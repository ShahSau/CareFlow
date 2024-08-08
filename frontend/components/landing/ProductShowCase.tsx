import Image from 'next/image'
import React from 'react'

const ProductShowCase = () => {
  return (
   <section className='bg-gradient-to-b from-[#bec6e1] to-[#6284f2] py-24 overflow-x-clip'>
         <div className='container'>
            <div className='max-w-[540px] mx-auto'>
                <div className='flex justify-center'>
                    <div className='text-sm inline-flex border border-[#000]/50 px-3 py-1 rounded-lg tracking-tight'>Health is the key</div>
                </div>

                <h2 className='text-center text-3xl md:text-[54px] md:leading-[60px] font-bold tracking-tighter bg-gradient-to-b from-[#6f8dea] to-[#011453] text-transparent bg-clip-text mt-6'>
                    Get the best treatment with upmost care.
                </h2>

                <p className='text-center text-[22px] leading-[30px] tracking-tight text-[#010D3E] mt-5'>
                    It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using
                </p>

            </div>
            <div className='relative flex justify-center'>
                <Image
                    src='/assets/images/product-image.png'
                    width={800}
                    height={600}
                    alt='Product Showcase'
                    className='mt-10'
                />
                <Image
                    src='/assets/images/pyramid.png'
                    width={262}
                    height={262}
                    alt='Pyramid Image'
                    className='hidden md:block absolute -right-8 -top-24'
                />
                <Image
                    src='/assets/images/tube.png'
                    width={248}
                    height={248}
                    alt='Tube Image'
                    className='hidden md:block absolute bottom-20 -left-10'
                />
            </div>
            
         </div>
   </section>
  )
}

export default ProductShowCase