<script context="module">
  export function preload({ params, query }) {
    return this.fetch(`blog.json`).then(r => r.json()).then(posts => {
      return { posts };
    });
  }
</script>

<script>
  import { siteUrl } from '../../stores/_config.js';

  export let posts;
</script>

<style>
  h2,
  .post-item-footer {
    font-family: Rubik, sans-serif;
    font-weight: 700;
  }

  .post-item-date {
    color: #AAA;
    text-align: left;
    text-transform: uppercase;
    margin-right: 16px;
  }

  hr {
    margin: 60px auto;
  }
</style>

<svelte:head>
  <title>Muktazam's Blog</title>
  <meta name="description" content="This blog is made with ❤️ with Sapper." />
  <meta name="keywords" content="profile, jamstack, frontend, backend, fullstack, ui, ux, ux designer, ui designer, " />

  <!-- Open Graph / Facebook -->
  <meta property="og:type" content="website">
  <meta property="og:url" content="{siteUrl}/blog/">
  <meta property="og:title" content="Muktazam's Blog">
  <meta property="og:description" content="This blog is made with ❤️ with Sapper.">
  <meta property="og:image" content="https://avatars3.githubusercontent.com/u/10528000?v=3&s=300">

  <!-- Twitter -->
  <meta property="twitter:card" content="summary_large_image">
  <meta property="twitter:url" content="{siteUrl}/blog/">
  <meta property="twitter:title" content="Muktazam's Blog">
  <meta property="twitter:description" content="This blog is made with ❤️ with Sapper.">
  <meta property="twitter:image" content="https://avatars3.githubusercontent.com/u/10528000?v=3&s=300">
</svelte:head>

<div class="container">
  <h1>Blog</h1>
  {#each posts as post, index}
    {#if index}
      <hr />
    {/if}
    <div class="post-item">
      <h2>
        <a rel='prefetch' href='blog/{post.slug}'>{post.title}</a>
      </h2>
      <p>{post.excerpt}</p>
      <div class="post-item-footer">
        <span class="post-item-date">— {post.printDate}</span>
      </div>
    </div>
  {/each}
</div>