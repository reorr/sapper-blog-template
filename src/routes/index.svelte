<script>
  import { siteUrl } from '../stores/_config.js';  
</script>

<style>
  .home-container {
    align-items: center;
    display: flex;
    flex: 1;
    justify-content: center;
    margin: 2em 0;
    min-height: 400px;
  }

  .home-copy {
    flex: 1;
  }

  h1 {
    font-weight: 700;
    margin-bottom: 0.5em;
  }

  p {
    font-size: 1.4em;
    line-height: 1.5;
  }

  figure {
    margin: 0 1em;
    text-align: center;
  }

  figcaption {
    font-size: .8em;
    font-style: italic;
  }

  img {
    width: 100%;
    max-width: 400px;
    border-radius: 50%;
  }

  @media (max-width: 1020px) {
    p {
      font-size: 1.2em;
    }

    img {
      max-width: 300px;
    }
  }

  @media (max-width: 800px) {
    .home-container {
      flex-direction: column;
    }

    .home-copy {
      flex: 0;
      padding-bottom: 2em;
      text-align: center;
    }
  }
</style>

<svelte:head>
  <title>Muktazam Hasbi Ashidiqi</title>
  <meta name="description"
    content="My name is Muktazam Hasbi. I writes about tech, mainly linux stuff and my stories in this blog. Recently interested with JAMstack technologies. This blog is made with ❤️ with Sapper." />
  <meta name="keywords" content="profile, jamstack, frontend, backend, fullstack, ui, ux, ux designer, ui designer, " />

  <!-- Open Graph / Facebook -->
  <meta property="og:type" content="website">
  <meta property="og:url" content="{siteUrl}/">
  <meta property="og:title" content="Muktazam Hasbi Ashidiqi">
  <meta property="og:description"
    content="My name is Muktazam Hasbi. I writes about tech, mainly linux stuff and my stories in this blog. Recently interested with JAMstack technologies. This blog is made with ❤️ with Sapper.">
  <meta property="og:image" content="https://avatars3.githubusercontent.com/u/10528000?v=3&s=300">

  <!-- Twitter -->
  <meta property="twitter:card" content="summary_large_image">
  <meta property="twitter:url" content="{siteUrl}/">
  <meta property="twitter:title" content="Muktazam Hasbi Ashidiqi">
  <meta property="twitter:description"
    content="My name is Muktazam Hasbi. I writes about tech, mainly linux stuff and my stories in this blog. Recently interested with JAMstack technologies. This blog is made with ❤️ with Sapper.">
  <meta property="twitter:image" content="https://avatars3.githubusercontent.com/u/10528000?v=3&s=300">
</svelte:head>

<div class="home-container">
  <div class="home-copy">
    <h1>Muktazam Hasbi Ashidiqi</h1>
    <p>Did you know that peaches, pears, apricots, quinces, strawberries, and apples are members of the rose family?
      They are come from the same <strong>Rosaceae</strong> family.</p>
  </div>

  <figure>
    <img alt='Profile photo from github' src='https://avatars3.githubusercontent.com/u/10528000?v=3&s=300'>
    <!-- <figcaption>Illustration thanks to <a href="https://undraw.co" target="_blank">Undraw</a></figcaption> -->
  </figure>
</div>