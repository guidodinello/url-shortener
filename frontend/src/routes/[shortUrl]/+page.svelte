<script>
    /**
     * @typedef {Object} Payload
     * @property {string} Url - The URL property in the payload object.
     */
    import { page } from "$app/stores";
    const shortUrl = $page.params.shortUrl; // $params contains the dynamic parameters

    import { onMount } from "svelte";

    // Function to fetch the long URL from the server
    /**
     * @param {string} shortUrl
     * @returns {Promise<string>}
     */
    async function fetchLongUrl(shortUrl) {
        /** @type {Payload} */
        const payload = { Url: shortUrl };
        const response = await fetch("/api/redirect", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(payload),
        });

        if (!response.ok) {
            throw new Error(`Failed to fetch: ${response.statusText}`);
        }

        const data = await response.json();
        return data.url;
    }

    /**
     * @type {string | null}
     */
    let longUrl = null;
    onMount(async () => {
        longUrl = await fetchLongUrl(shortUrl);
        window.location.href = longUrl ?? window.location.href;
    });
</script>

{#if longUrl === null}
    <div class="spinner">
        <div class="double-bounce1"></div>
        <div class="double-bounce2"></div>
    </div>
    <p>Loading...</p>
{:else}
    <p>Redirecting to: {longUrl}</p>
{/if}

<style>
    .spinner {
        width: 40px;
        height: 40px;
        position: relative;
        margin: 100px auto;
    }

    .double-bounce1,
    .double-bounce2 {
        width: 100%;
        height: 100%;
        border-radius: 50%;
        background-color: #333;
        opacity: 0.6;
        position: absolute;
        top: 0;
        left: 0;
        -webkit-animation: sk-bounce 2s infinite ease-in-out;
        animation: sk-bounce 2s infinite ease-in-out;
    }
</style>
