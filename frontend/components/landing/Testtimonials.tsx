import Image from "next/image";
import { twMerge } from "tailwind-merge";



const testimonials = [
  {
    text: "As a seasoned designer always on the lookout for innovative tools, Framer.com instantly grabbed my attention.",
    imageSrc: '/assets/images/avatar-1.png',
    name: "Jamie Rivera",
    username: "@jamietechguru00",
  },
  {
    text: "Our team's productivity has skyrocketed since we started using this tool. ",
    imageSrc: '/assets/images/avatar-2.png',
    name: "Josh Smith",
    username: "@jjsmith",
  },
  {
    text: "This app has completely transformed how I manage my projects and deadlines.",
    imageSrc: '/assets/images/avatar-3.png',
    name: "Morgan Lee",
    username: "@morganleewhiz",
  },
  {
    text: "I was amazed at how quickly we were able to integrate this app into our workflow.",
    imageSrc: '/assets/images/avatar-4.png',
    name: "Casey Jordan",
    username: "@caseyj",
  },
  {
    text: "Planning and executing events has never been easier. This app helps me keep track of all the moving parts, ensuring nothing slips through the cracks.",
    imageSrc: '/assets/images/avatar-5.png',
    name: "Taylor Kim",
    username: "@taylorkimm",
  },
  {
    text: "The customizability and integration capabilities of this app are top-notch.",
    imageSrc: '/assets/images/avatar-6.png',
    name: "Riley Smith",
    username: "@rileysmith1",
  },
  {
    text: "Adopting this app for our team has streamlined our project management and improved communication across the board.",
    imageSrc: '/assets/images/avatar-7.png',
    name: "Jordan Patels",
    username: "@jpatelsdesign",
  },
  {
    text: "With this app, we can easily assign tasks, track progress, and manage documents all in one place.",
    imageSrc: '/assets/images/avatar-8.png',
    name: "Sam Dawson",
    username: "@dawsontechtips",
  },
  {
    text: "Its user-friendly interface and robust features support our diverse needs.",
    imageSrc: '/assets/images/avatar-9.png',
    name: "Casey Harper",
    username: "@casey09",
  },
];

const firstColumn = testimonials.slice(0, 3);
const secondColumn = testimonials.slice(3, 6);
const thirdColumn = testimonials.slice(6, 9);

const TestimonialsColumn = (props:{className?:string; testimonials: typeof testimonials}) => {
    return (
        <div className={twMerge("flex flex-col gap-6 mt-10 [mask-image:linear-gradient(to_bottom,transparent,black_25%,black_75%,transparent)]", props.className)}>
                    {props.testimonials.map(({ text, imageSrc, name, username }, index) => (
                        <div key={index} className="p-10 border border-solid border-[#222222]/10 rounded-3xl shadow-[0_7px_14px_#EAEAEA] max-w-xs w-full">
                            <div>
                                {text}
                            </div>
                            <div className="flex items-center gap-2 mt-5">
                                <Image 
                                    src={imageSrc}
                                    alt={name} 
                                    width={40} 
                                    height={40} 
                                    className="rounded-full"
                                />
                                <div className="flex flex-col">
                                    <div className="font-medium tracking-tight leading-5">{name}</div>
                                    <p className="leading-5 tracking-tight">{username}</p>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
    )}


export const Testimonials = () => {
  return (
    <section className="bg-[#6284f2]">
        <div className="container">
            <div className="max-w-[540px] mx-auto">
                <div className="flex justify-center">
                <div className="text-sm inline-flex border border-[#000]/50 px-3 py-1 rounded-lg tracking-tight">Testimonials</div>
                </div>
                <h2 className="text-center text-3xl md:text-[54px] md:leading-[60px] font-bold tracking-tighter bg-gradient-to-b from-[#6f8dea] to-[#011453] text-transparent bg-clip-text mt-6">What are user say</h2>
                <p className="text-center text-[22px] leading-[30px] tracking-tight text-[#010D3E] mt-5">
                Various versions have evolved over the years, sometimes by accident, sometimes on purpose injected humour and the like.
                </p>
            </div>
            <div className="flex justify-center gap-6">
                <TestimonialsColumn testimonials={firstColumn}/>
                <TestimonialsColumn testimonials={secondColumn}  className="hidden md:flex"/>
                <TestimonialsColumn testimonials={thirdColumn}  className="hidden lg:flex"/>
            </div>
            
        </div>
    </section>
  )};
