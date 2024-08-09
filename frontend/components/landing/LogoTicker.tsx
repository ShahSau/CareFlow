"use client";

import Image from 'next/image'
import React from 'react'
import { motion } from 'framer-motion'

const LogoTicker = () => {
  return (
    <div className='bg-[#bec6e1] py-8 md:py-12'>
        <div className='container'>
            <div className='flex overflow-hidden'>
                <motion.div 
                    className='flex gap-14 flex-none pr-14'
                    animate={{
                        translateX:"-50%",
                    }}
                    transition={{
                        duration:20,
                        repeat:Infinity,
                        repeatType:'loop',
                        ease:'linear'
                    }}
                >
                    <Image
                        src='/assets/images/logo-acme.png'
                        height={32}
                        width={32}
                        alt='Acme Logo'
                        className='logo-ticker-image'
                    />
                    <Image
                        src='/assets/images/logo-apex.png'
                        height={32}
                        width={32}
                        alt='Apex Logo'
                        className='logo-ticker-image'
                    />
                    <Image
                        src='/assets/images/logo-celestial.png'
                        height={32}
                        width={32}
                        alt='Celestial Logo'
                        className='logo-ticker-image'
                    />
                    <Image
                        src='/assets/images/logo-echo.png'
                        height={32}
                        width={32}
                        alt='Echo Logo'
                        className='logo-ticker-image'
                    />
                    <Image
                        src='/assets/images/logo-pulse.png'
                        height={32}
                        width={32}
                        alt='Pulse Logo'
                        className='logo-ticker-image'
                    />
                    <Image
                        src='/assets/images/logo-quantum.png'
                        height={32}
                        width={32}
                        alt='Quantum Logo'
                        className='logo-ticker-image'
                    />

                    {/* second set of logos for animation*/}

                    <Image
                        src='/assets/images/logo-acme.png'
                        height={32}
                        width={32}
                        alt='Acme Logo'
                        className='logo-ticker-image'
                    />
                    <Image
                        src='/assets/images/logo-apex.png'
                        height={32}
                        width={32}
                        alt='Apex Logo'
                        className='logo-ticker-image'
                    />
                    <Image
                        src='/assets/images/logo-celestial.png'
                        height={32}
                        width={32}
                        alt='Celestial Logo'
                        className='logo-ticker-image'
                    />
                    <Image
                        src='/assets/images/logo-echo.png'
                        height={32}
                        width={32}
                        alt='Echo Logo'
                        className='logo-ticker-image'
                    />
                    <Image
                        src='/assets/images/logo-pulse.png'
                        height={32}
                        width={32}
                        alt='Pulse Logo'
                        className='logo-ticker-image'
                    />
                    <Image
                        src='/assets/images/logo-quantum.png'
                        height={32}
                        width={32}
                        alt='Quantum Logo'
                        className='logo-ticker-image'
                    />
                </motion.div>
            </div>

        </div>

    </div>
  )
}

export default LogoTicker