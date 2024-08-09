import CallToAction from "@/components/landing/CallToAction";
import Header from "@/components/landing/Header";
import Hero from "@/components/landing/Hero";
import LogoTicker from "@/components/landing/LogoTicker";
import ProductShowCase from "@/components/landing/ProductShowCase";
import { Testimonials } from "@/components/landing/Testtimonials";


export default function Home() {
  return (
    <>
      <Header />
      <Hero />
      <LogoTicker />
      <ProductShowCase />
      <Testimonials />
      <CallToAction />
    </>
 
  );
}
