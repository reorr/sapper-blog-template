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
  <title>Sapper Blog Template</title>
</svelte:head>

<div class="home-container">
  <div class="home-copy">
    <h1>Muktazam Hasbi Ashidiqi</h1>
    <p>Did you know that peaches, pears, apricots, quinces, strawberries, and apples are members of the rose family?</p>
    <p>They are come from the same <strong>Rosaceae</strong> family.</p>
  </div>

  <figure>
    <img alt='Profile photo from github' src='https://avatars3.githubusercontent.com/u/10528000?v=3&s=300'>
    <!-- <figcaption>Illustration thanks to <a href="https://undraw.co" target="_blank">Undraw</a></figcaption> -->
  </figure>
</div>