<script>
    /**
     * @typedef {Object} Payload
     * @property {string} Url - The URL property in the payload object.
     */

    import { afterUpdate, onMount } from "svelte";

    onMount(() => {
        document
            .querySelector("form")
            ?.addEventListener("submit", handleSubmit);
    });

    /**
     * @type {HTMLInputElement}
     */
    let shortUrlInput;
    // Add the focus class to shortUrlInput when shortUrl changes
    $: if (shortUrlInput && shortUrl) {
        shortUrlInput.select();
    }
    // Remove the focus class from shortUrlInput after each update
    afterUpdate(() => {
        if (shortUrlInput) {
            shortUrlInput.classList.remove("focus");
        }
    });

    /**
     * @type {string | undefined}
     */
    let shortUrl = "";

    const toClipboard = () => {
        if (!shortUrl) return;
        navigator.clipboard.writeText(shortUrl);
        // more feedback, tick icon or something
    };

    /**
     * @param {Event} e
     */
    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            /**
             * @type {Payload}
             */
            const payload = {
                // @ts-ignore
                Url: e.target?.url.value,
            };
            const response = await fetch("/api/shorten", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(payload),
            });

            console.log(response);

            if (response.ok) {
                const data = await response.json();
                shortUrl = `${import.meta.env.VITE_FRONT_BASE_URL}/${
                    data.shortUrl
                }`;
            } else {
                console.error("Error:", response.statusText);
            }
        } catch (error) {
            console.error("Error:", error);
        }
    };
</script>

<h1>URL Shortener</h1>

<form class="row mt-5">
    <div class="col-75 input">
        <label for="url"></label>
        <input
            type="url"
            id="url"
            name="url"
            placeholder="Super long URL: https://www.this-is-a-really-long-url.com"
        />
    </div>
    <div class="col-25 float-end">
        <button type="submit">Shorten</button>
    </div>
</form>

<section class="row mt-5 mb-2">
    <div class="col-75 input">
        <label for="short-url"></label>
        <input
            bind:this={shortUrlInput}
            type="url"
            id="short-url"
            name="short-url"
            placeholder="Short URL &#128522"
            value={shortUrl}
        />
    </div>
    <div class="col-25 float-end">
        <button on:click={toClipboard} disabled={!shortUrl}>Copy</button>
    </div>
</section>

<style>
    button {
        background-color: #4caf50;
        border: none;
        color: white;
        padding: 16px 30px;
        text-align: center;
        font-size: 16px;
        margin: 4px 2px;
        transition: 0.3s;
        display: inline-block;
        text-decoration: none;
        cursor: pointer;
        border-radius: 10px;
        font-weight: bold;

        width: 75%;

        letter-spacing: 0.08rem;
    }

    button:disabled {
        background-color: rgba(81, 199, 85, 0.9);
        color: lightgray;
        cursor: not-allowed;
    }

    button:hover {
        background-color: #3e8e41;
        scale: calc(1.01);
    }

    input[type="url"] {
        background: #fff;
        border: 1px solid #d3d4d7;
        border-radius: 8px;
        padding: 16px 20px;
        font-size: 20px;
        line-height: 24px;
        height: auto;
        margin-top: 8px;
        margin-bottom: 0;
        box-shadow: none;
    }

    input[type="url"]:hover {
        outline: none;
    }
</style>
