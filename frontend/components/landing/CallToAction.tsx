import React from 'react'
import { Button } from '../ui/button'
import Image from 'next/image'

const CallToAction = () => {
  return (
    <section className='bg-gradient-to-b from-[#6284f2] to-[#bec6e1] py-24 overflow-x-clip'>
        <div className='container'>
            <div className='max-w-[540px] mx-auto relative'>
                <h2 className='text-center text-3xl md:text-[54px] md:leading-[60px] font-bold tracking-tighter bg-gradient-to-b from-[#6f8dea] to-[#011453] text-transparent bg-clip-text mt-6'>Sing up for free now</h2>
                <p className='text-center text-[22px] leading-[30px] tracking-tight text-[#010D3E] mt-5'>
                    It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing
                </p>
                <Image
                    src='/assets/images/star.png'
                    width={260}
                    height={260}
                    alt='Star Image'
                    className='absolute -left-[250px] -top-[30px] hidden md:block'
                />
                <Image
                    src='/assets/images/spring.png'
                    width={260}
                    height={260}
                    alt='Spring Image'
                    className='absolute -right-[250px] top-[10px] hidden md:block'
                />
            </div>
            <div className='flex gap-2 mt-10 justify-center'>
               <Button className='shad-primary-btn'>
                    Book an appointment
                </Button>
                <Button variant="outline">
                    Join us as a doctor
                </Button>
            </div>
        </div>
    </section>
  )
}

export default CallToAction