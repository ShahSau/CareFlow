import Image from 'next/image'
import React from 'react'

const Header = () => {
  return (
    <header className='sticky top-0 backdrop-blur-sm z-20'>
        <div className='flex justify-center items-center py-3 bg-slate-300 text-sm gap-3'>
            <p className='text-black/60 hidden md:block'>Get the best treatment with upmost care.</p>
            <div className='inline-flex gap-1 items-center'>
                <p className='text-slate-600'>
                    Get started with CareFlow
                </p>
                <Image 
                    src="/assets/images/arrow-right.svg" 
                    alt="CareFlow" 
                    className='inline-flex justify-center'
                    width={20}
                    height={20}
                />
            </div>
        </div>
        <div className='py-5'>
            <div className='container'>
                <div className='flex items-center justify-between'> 
                    <Image
                        src="/assets/icons/logo-icon.svg"
                        alt="CareFlow"
                        width={40}
                        height={40}
                    />
                    <Image
                        src="/assets/icons/menu.svg"
                        alt="CareFlow"
                        width={20}
                        height={20}
                        className='bg-white md:hidden'
                    />
                    <nav className='hidden md:flex gap-6 text-white/60 items-center'>
                            <a href='#' className=''>Appoinment</a>
                            <a href='#' className=''>Join Us</a>
                            
                    </nav>

                </div>
                
            </div>
            
        </div>
    </header>
  )
}

export default Header