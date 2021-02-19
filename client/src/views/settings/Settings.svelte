<script>
  import DebtRecap from "./DebtRecap.svelte";
  import Reset from "./Reset.svelte";
  import PriceSection from "./PriceSection.svelte";
  import Container from "../../components/Container.svelte";
  import Watermark from "./Watermark.svelte";
</script>

<Container>
  <PriceSection />
  <DebtRecap />
  <Reset />
  <Watermark />
</Container>
