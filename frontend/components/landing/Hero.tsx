"use client";

import React, { useRef } from 'react'
import { Button } from '../ui/button'
import {motion, useScroll, useTransform} from 'framer-motion'
import Link from 'next/link';


const Hero = () => {
    const heroRef = useRef(null);
    const { scrollYProgress } = useScroll({
        target: heroRef,
        offset:["start end", "end start"],
    });
    const translateY = useTransform(scrollYProgress, [0, 1], [150, -150]);


  return (
    <section ref={heroRef} className=' pt-8 pb-20 md:pt-5 md:pb-10 bg-[radial-gradient(ellipse_200%_100%_at_bottom_left,#183EC2,#131619_70%)] overflow-x-clip'>
        <div className='container overflow-hidden'>
            <div className='md:flex items-center'>
                <div className='md:w-[478px]'>
                    <div className='text-sm inline-flex border border-[#f0f0f0]/20 px-3 py-1 rounded-lg tracking-tight'>Better than before</div>
                    <h1 className='text-5xl md:text-7xl font-bold tracking-tighter bg-gradient-to-b from-[#f0f0f0] to-[#001E80] text-transparent bg-clip-text mt-6'>Get the best treatment with upmost care.</h1>
                    <p className='text-xl text-[#f0f0f0] tracking-tight mt-6'>It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using </p>
                    <div className='flex gap-1 items-center mt-[30px]'>
                        <Button className='shad-primary-btn'>
                            <Link href='/patients'>
                                Book an appointment
                            </Link>
                        </Button>
                        <Button variant="outline">
                            <Link href='/doctors'>
                                Join us as a doctor
                            </Link>
                        </Button>
                    </div>
                </div>
                <div className='mt-20 md:mt-0 md:h-[648px] md:flex-1 relative overflow-hidden'>
                    <motion.img 
                        src='/assets/images/cog.png' 
                        width={200} height={200} 
                        alt='Cog Image' 
                        className='md:absolute md:h-full md:w-auto md:max-w-none md:-left-6 lg:left-0'
                        // animate={{
                        //     translateY:[-30,30]
                        // }}
                        // transition={{
                        //     duration:3,
                        //     repeat:Infinity,
                        //     repeatType:'mirror',
                        //     ease:'easeInOut'
                        // }}
                        initial={{ rotate: 0 }}
                        animate={{ rotate: 60 }}
                        transition={{ 
                            duration: 2, 
                            repeat: Infinity, 
                            ease: "linear",
                            repeatType:'mirror'
                        }}

                    />
                    <motion.img
                        src='/assets/images/cylinder.png'
                        width={220}
                        height={220}
                        alt='Cylinder Image'
                        className='hidden md:block -top-4 -left-4 md:absolute'
                        style={{
                            translateY
                        }}
                    />
                    <motion.img
                        src='/assets/images/noodle.png'
                        width={200}
                        height={200}
                        alt='Noodle Image'
                        className='hidden lg:block absolute top-[400px] left-[550px]'
                        style={{
                            rotate:30,
                            translateY
                        }}
                    />
                </div>
            </div>
            
        </div>
    </section>
  )
}

export default Hero