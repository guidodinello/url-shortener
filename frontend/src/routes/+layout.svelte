<script>
    import "../styles/global.css";
</script>

<nav>
    <section class="logo">
        <a href="/">
            <div class="logo-img">
                <img
                    src="https://assets-global.website-files.com/60f1e0a61b660f482182420a/60f85a80e4bfc21507c286d3_Shorty%27s%20Logo-min-p-1600.png"
                    alt="logo"
                />
            </div>
        </a>
    </section>

    <section class="right-menu float-end">
        <a href="/about">About</a>
        <a href="/contact">Contact</a>
        <a href="/register">Register</a>
    </section>
</nav>

<main id="main-container">
    <section id="menu">
        <slot />
    </section>
</main>

<style>
    nav {
        display: flex;
        justify-content: space-around;
        align-items: center;
        padding: 1rem;
    }

    nav a {
        text-decoration: none;
    }

    nav a:hover {
        text-decoration: underline;
    }

    #main-container {
        display: flex;
        justify-content: center;

        width: 100vw;
        height: 100vh;
    }

    #menu {
        box-shadow: none;
        border: 3px #e8e9eb solid;
        border-radius: 16px;

        padding: 2rem 3rem;
        margin-top: 5%;

        height: fit-content;
        width: 66vw;
        max-height: 66vh;
        max-width: 800px;

        background-color: #fff;
    }

    .right-menu {
        display: flex;
        gap: 1.5rem;
        font-weight: 400;
        align-items: center;
        width: 30%;
        text-decoration: none;
    }

    .right-menu a {
        text-decoration: none;
        color: black;
    }

    .logo-img {
        width: 120px;
        height: 80px;
        border-radius: 50%;
        background-color: black;
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .logo img {
        width: 100px;
    }
</style>
