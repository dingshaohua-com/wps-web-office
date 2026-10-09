import WebOfficeSDK from '@/utils/web-office-sdk-solution-v2.0.7/web-office-sdk-solution-v2.0.7.es.js';
import { useEffect, useRef } from 'react';


// const fileUrl = 'https://dsh-1300009960.cos.ap-beijing.myqcloud.com/office/jl.docx'

// const instance = WebOfficeSDK.init({
//   officeType: WebOfficeSDK.OfficeType.Writer,
//   appId: 'SX20260914ZIHJLM',
//   // fileId: 'HyIKDNdzqFvYFCsmkfQOXAGVSDOCvMTk'
//   fileId:'jl_docx',
//   mount: document.querySelector('#wps-container')
// });

// console.log(instance);



function About() {

  const containerRef = useRef<HTMLDivElement>(null)
  useEffect(() => {
    if (!containerRef.current) return

    const instance = WebOfficeSDK.init({
      officeType: WebOfficeSDK.OfficeType.Writer,
      appId: 'SX20260914ZIHJLM',
      fileId: 'jl_docx',
      mount: containerRef.current
    })

    instance.on('fileOpen', (result: unknown) => {
      console.log('fileOpen:', result)
    })

    instance.on('error', (error: unknown) => {
      console.error('WebOffice error:', error)
    })

    return () => {
      instance.destroy()
    }
  }, [])

  return (
    <div ref={containerRef} style={{ width: '100%', height: '100vh' }}></div>
  )
}

export default About
