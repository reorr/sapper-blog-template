<script>
  import { siteUrl } from '../stores/_config.js';  
</script>

<style>
  figure {
    float: right;
    margin: -3em -7em 2em 2em;
    max-width: 400px;
  }

  figure img {
    border-radius: 8px;
  }

  @media (max-width: 1020px) {
    figure {
      float: none;
      margin: 0 auto 2em;
    }
  }
</style>

<svelte:head>
  <title>About</title>
  <meta name="description"
    content="My name is Muktazam Hasbi. I writes about tech, mainly linux stuff and my stories in this blog. Recently interested with JAMstack technologies. This blog is made with ❤️ with Sapper." />
  <meta name="keywords" content="profile, jamstack, frontend, backend, fullstack, ui, ux, ux designer, ui designer, " />

  <!-- Open Graph / Facebook -->
  <meta property="og:type" content="website">
  <meta property="og:url" content="{siteUrl}/about">
  <meta property="og:title" content="Muktazam Hasbi Ashidiqi">
  <meta property="og:description"
    content="My name is Muktazam Hasbi. I writes about tech, mainly linux stuff and my stories in this blog. Recently interested with JAMstack technologies. This blog is made with ❤️ with Sapper.">
  <meta property="og:image" content="{siteUrl}/rsz_florian-klauer-489-unsplash.jpg">

  <!-- Twitter -->
  <meta property="twitter:card" content="summary_large_image">
  <meta property="twitter:url" content="{siteUrl}/about">
  <meta property="twitter:title" content="Muktazam Hasbi Ashidiqi">
  <meta property="twitter:description"
    content="My name is Muktazam Hasbi. I writes about tech, mainly linux stuff and my stories in this blog. Recently interested with JAMstack technologies. This blog is made with ❤️ with Sapper.">
  <meta property="twitter:image" content="{siteUrl}/rsz_florian-klauer-489-unsplash.jpg">
</svelte:head>

<div class="container">
  <h1>About</h1>
  <figure>
    <img src='rsz_florian-klauer-489-unsplash.jpg' alt='Image of a vintage typewriter.'>
    <figcaption>Photo by <a href="https://unsplash.com/@florianklauer" target="_blank">Florian Klauer</a> on Unsplash
    </figcaption>
  </figure>
  <p>Text placeholder via <a href="https://jeffsum.com/" target="_blank">Jeffsum</a>.</p>
  <p>So you two dig up, dig up dinosaurs? What do they got in there? King Kong? My dad once told me, laugh and the world
    laughs with you, Cry, and I'll give you something to cry about you little bastard! Life finds a way. God creates
    dinosaurs. God destroys dinosaurs. God creates Man. Man destroys God. Man creates Dinosaurs.</p>
  <p>You really think you can fly that thing? You know what? It is beets. I've crashed into a beet truck. Forget the fat
    lady! You're obsessed with the fat lady! Drive us out of here! Is this my espresso machine? Wh-what is-h-how did you
    get my espresso machine?</p>
  <p>Hey, you know how I'm, like, always trying to save the planet? Here's my chance. Hey, take a look at the
    earthlings. Goodbye! I was part of something special. Just my luck, no ice. You're a very talented young man, with
    your own clever thoughts and ideas. Do you need a manager?</p>
  <p>Jaguar shark! So tell me - does it really exist? This thing comes fully loaded. AM/FM radio, reclining bucket
    seats, and... power windows. Yes, Yes, without the oops! You're a very talented young man, with your own clever
    thoughts and ideas. Do you need a manager?</p>
  <p>Yes, Yes, without the oops! Do you have any idea how long it takes those cups to decompose. They're using our own
    satellites against us. And the clock is ticking. Do you have any idea how long it takes those cups to decompose. My
    dad once told me, laugh and the world laughs with you, Cry, and I'll give you something to cry about you little
    bastard!</p>
</div>