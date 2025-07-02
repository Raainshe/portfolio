let info = {
  name: "Ryan Makoni",
  logo_name: "makoni",
  flat_picture: require("./src/assets/ryan.jpg"),
  config: {
    use_cookies: true,
    navbar: {
      blur: false
    }
  },
  description:
    "Hey there! I'm Ryan, a passionate Full Stack Developer who loves turning ideas into reality through code. I'm currently pursuing Software Engineering at 42 Heilbronn in Germany, where I'm diving deep into project-based learning and peer-to-peer collaboration.<br><br>With over 5 years of experience under my belt, I've had the privilege of wearing many hats in the tech world. I founded and run Key Space, a technology company focused on web development services, where I lead projects from concept to deployment. I've also worked as a Web Developer at Tiny Optics, where I handle full-stack development and API integrations, and as an IT Facilitator at CTU Training Solutions, where I get to share my knowledge and mentor aspiring developers.<br><br>I'm passionate about modern web technologies like React, Vue.js, ASP.NET, and cloud platforms like Microsoft Azure. What drives me is staying current with emerging technologies and delivering innovative solutions that make a real impact. I believe in continuous learning and pushing the boundaries of what's possible in web development.",
  links: {
    linkedin: "https://www.linkedin.com/in/ryan-makoni/",
    github: "https://github.com/Raainshe",
    angellist: "https://angel.co/u/ryan-makoni",
    resume: "/RyanMakoniResume.pdf"
  },
  education: [
    {
      name: "42 Heilbronn",
      place: "Heilbronn, Germany",
      date: "Oct, 2024 - Present",
      degree: "Software Engineering",
      gpa: "Project-based learning",
      description:
        "Project-based learning school, based on peer-to-peer learning. The 42 program takes a project-based approach to progress and is designed to develop technical and people skills that match the expectations of the labour market.",
      skills: [
        "Software Engineering",
        "Project-based Learning",
        "Peer-to-Peer Learning",
        "Technical Skills",
        "Problem Solving"
      ]
    },
    {
      name: "Richfield Graduate Institute of Technology",
      place: "South Africa",
      date: "2017 - 2020",
      degree: "Bachelor of Science - Information Technology",
      gpa: "Specialisation: Systems Development",
      description:
        "Graduated with a Bachelor of Science in Information Technology, specializing in Systems Development. This program provided a strong foundation in software development, database management, and IT systems.",
      skills: [
        "Information Technology",
        "Systems Development",
        "Software Development",
        "Database Management",
        "IT Systems"
      ]
    }
  ],
  experience: [
    {
      name: "Key Space",
      place: "Randburg, South Africa",
      date: "2023 - Oct, 2024",
      position: "Director",
      description:
        "Founded and manage a technology company focusing on web development services. Lead development projects for multiple clients, from concept to deployment. Oversee project timelines, client communications, and technical implementation.",
      skills: ["Project Management", "Web Development", "Client Relations", "Business Strategy"]
    },
    {
      name: "Tiny Optics",
      place: "Cape Town, South Africa",
      date: "2020 - Present",
      position: "Web Developer",
      description:
        "Performed full-stack development duties for client websites. Implemented and integrated various APIs. Conducted testing, bug fixes, and regular content updates.",
      skills: ["Full Stack Development", "API Integration", "Testing", "Bug Fixes", "Project Management"]
    },
    {
      name: "CTU Training Solutions",
      place: "Randburg, South Africa",
      date: "2022 - 2024",
      position: "IT Facilitator",
      description:
        "Taught UI design, including HTML, JavaScript, Java, and Adobe Suite. Developed curriculum materials and student assessments. Provided mentoring and technical guidance to aspiring developers.",
      skills: ["Teaching", "Curriculum Development", "HTML", "JavaScript", "Python", "Adobe Suite"]
    },
    {
      name: "Value Logistics",
      place: "Germiston, South Africa",
      date: "2022",
      position: "Junior Developer",
      description:
        "Improved user engagement by 25% by developing high-performance web applications with ASP.NET and MVC. Created custom SQL procedures and reports for business intelligence. Implemented database solutions for enterprise applications.",
      skills: ["ASP.NET", "MVC", "SQL", "Business Intelligence", "Database Solutions"]
    },
    {
      name: "Value Logistics",
      place: "Germiston, South Africa",
      date: "2021 - 2022",
      position: "Web Development Intern",
      description:
        "Generated client reports using SQL and SSRS. Developed custom web applications for internal business needs. Maintained and hosted web applications for various departments.",
      skills: ["SQL", "SSRS", "Web Applications", "Application Hosting"]
    },
    {
      name: "Freelance Web Developer",
      place: "Randburg, South Africa",
      date: "2019 - 2022",
      position: "Freelance Web Developer",
      description:
        "Designed and developed websites ranging from basic informational sites to complex e-commerce platforms. Created custom environments for clients' internal use. Provided ongoing maintenance and hosting services.",
      skills: ["E-commerce", "Website Design", "Custom Development", "Maintenance", "Hosting"]
    },
    {
      name: "Master Your Finances",
      place: "Randburg, South Africa",
      date: "2019 - 2021",
      position: "Web Developer & IT Technician",
      description:
        "Developed and maintained the company website. Provided technical support for livestreams and virtual meetings. Created digital marketing materials and edited promotional content.",
      skills: ["Website Maintenance", "Technical Support", "Digital Marketing", "Content Creation"]
    }
  ],
  skills: [
    {
      title: "programming languages",
      info: [
        "JavaScript",
        "TypeScript",
        "C#",
        "C",
        "C++",
        "Java",
        "Kotlin",
        "HTML",
        "CSS",
        "Python"
      ],
      icon: "fa fa-code"
    },
    {
      title: "web technologies",
      info: [
        "React",
        "Vue.js",
        "ASP.NET",
        "MVC",
        "WordPress",
        "E-commerce",
        "Node.js",
        "Express.js",
        "Next.js",
        "Bootstrap",
        "Material UI",
      ],
      icon: "fas fa-laptop-code"
    },
    {
      title: "databases & reporting",
      info: [
        "SQL",
        "MongoDB",
        "SSRS"
      ],
      icon: "fa fa-database"
    },
    {
      title: "cloud & tools",
      info: [
        "Microsoft Azure",
        "GitHub",
        "P2P"
      ],
      icon: "fas fa-tools"
    },
    {
      title: "design & creative",
      info: [
        "Adobe CC",
        "UI Design",
        "Figma",
        "Premiere Pro",
      ],
      icon: "fa fa-pencil-square-o"
    },
    {
      title: "certifications",
      info: [
        "Microsoft Certified: Azure Fundamentals",
        "Microsoft Certified: DevOps Engineer Expert",
        "Microsoft Certified: Developer Associate"
      ],
      icon: "fa fa-certificate"
    }
  ],
  portfolio: [
    {
      name: "SBJ Application Form",
      pictures: [
        {
          img: require("./src/assets/SBJ.png")
        }
      ],
      technologies: ["Node.js", "Express.js", "Vue.js 3", "MongoDB", "Mongoose", "Passport.js", "SendGrid", "Google Drive API", "Trello API", "Bootstrap 5", "Pinia", "Axios", "Multer", "Heroku"],
      category: "Full Stack Application",
      date: "2023",
      github: "https://github.com/Tiny-Optics/SBJ-application-form",
      visit: "",
      description:
        "A comprehensive loan application form system with integration to Run A Loan API. Features include user authentication, file uploads, digital signatures, automated data synchronization, Google Drive integration for document storage, and Trello integration for workflow management. Built with Vue.js frontend and Node.js backend, featuring real-time form validation, email notifications via SendGrid, and scheduled API updates."
    },
    {
      name: "BTEC Solution",
      pictures: [
        {
          img: require("./src/assets/btecsolutions.png")
        },
        {
          img: require("./src/assets/btecsolutions2.png")
        }
      ],
      technologies: ["HTML", "CSS", "JavaScript", "ASP.NET", "C#"],
      category: "Web Application",
      date: "2023",
      github: "https://github.com/Raainshe/BTECSolution",
      visit: "https://btecsolutions.co.za/",
      description:
        "A web application built with ASP.NET and C# backend, featuring JavaScript and CSS frontend. The project demonstrates full-stack development capabilities with modern web technologies, providing a comprehensive solution for BTEC-related functionality."
    },
    {
      name: "Tuso Wellness (Landing Page)",
      pictures: [
        {
          img: require("./src/assets/Tuso.png")
        },
        {
          img: require("./src/assets/tuso2.png")
        }
      ],
      technologies: ["HTML", "CSS", "JavaScript", "Responsive Design", "ASP.NET"],
      category: "Landing Page",
      date: "2023",
      github: "https://github.com/Raainshe/TusoWellness",
      visit: "https://www.tusowellness.com/",
      description:
        "A modern, responsive landing page for Tuso Wellness. Designed and developed with clean, user-friendly interface focusing on wellness services. Features responsive design, smooth animations, and optimized for all devices."
    },
    {
      name: "Master Your Finances",
      pictures: [
        {
          img: require("./src/assets/masteryourfinances.png")
        },
        {
          img: require("./src/assets/masteryourfinances2.png")
        }
      ],
      technologies: ["HTML", "CSS", "JavaScript", "ASP.NET"],
      category: "Landing Page",
      date: "2019 - 2021",
      github: "",
      visit: "http://www.masteryourfinances.net/",
      description:
        "Developed and maintained the company website for Master Your Finances. Created digital marketing materials and provided ongoing technical support. The website serves as a platform for financial education and services."
    },
    {
      name: "Addaxz (E-commerce)",
      pictures: [
        {
          img: require("./src/assets/Addaxz.png")
        }
      ],
      technologies: ["HTML", "CSS", "JavaScript", "ASP.NET"],
      category: "E-commerce",
      date: "2022",
      github: "https://github.com/MbusoMakitla/AddaxZ",
      visit: "http://www.addaxz.co.za/",
      description:
        "A comprehensive e-commerce platform for Addaxz built from the ground up using ASP.NET. Features product catalog, shopping cart functionality, secure payment processing, and responsive design for optimal user experience. Includes an admin area where site administrators can manage products, orders, and view comprehensive site statistics."
    },
    {
      name: "Digikraal (E-commerce)",
      pictures: [
        {
          img: require("./src/assets/digikraal.png")
        },
        {
          img: require("./src/assets/digikraal2.png")
        }
      ],
      technologies: ["WordPress", "WooCommerce", "E-commerce", "HTML", "CSS", "JavaScript"],
      category: "E-commerce",
      date: "2022",
      github: "",
      visit: "http://www.digikraal.co.za",
      description:
        "A full-featured e-commerce website for Digikraal developed with WordPress and WooCommerce. Features product management, user authentication, payment integration, and mobile-responsive design for seamless shopping experience."
    },
    
  ],
  portfolio_design: [
    {
      name: "Emby-Torrent-Bot",
      title: "qBittorrent Discord Bot",
      pictures: [
        {
          img: require("./src/assets/emby torrentor meta ai.jpg")
        },
        {
          img: require("./src/assets/emby.png")
        }
      ],
      technologies: ["TypeScript", "Discord.js", "qBittorrent API", "Bun", "Node.js", "Chart.js"],
      category: "Discord Bot",
      github: "https://github.com/Raainshe/Emby-Torrent-Bot",
      date: "2024",
      visit: "",
      description:
        "A comprehensive Discord bot that allows users to interact with qBittorrent through Discord slash commands. Features include listing torrents, adding magnet links with category-based save paths, displaying download progress, deleting torrents, showing disk space usage with pie charts, and automatic seeding time management.<br/><br/>The bot includes intelligent seeding management that automatically stops torrents after 10x their download duration, helping maintain good seeding ratios while preventing indefinite seeding. Built with TypeScript and Discord.js, it provides a user-friendly interface for managing torrent downloads remotely through Discord commands."
    },
    {
      name: "cub3d",
      pictures: [
        {
          img: require("./src/assets/cub3d.png")
        },
        {
          img: require("./src/assets/cub3d2.png")
        }
      ],
      technologies: ["C", "Makefile", "Raycasting"],
      category: "Graphics/Game Engine",
      github: "https://github.com/Raainshe/cub3d",
      date: "2025",
      visit: "",
      description:
        "A raycasting 3D engine written in C, inspired by classic games like Wolfenstein 3D. This project demonstrates low-level graphics programming, real-time rendering, and game engine fundamentals. Features include map parsing, player movement, wall rendering, and basic texture mapping. Built as part of the 42 curriculum, it showcases strong skills in C, algorithms, and graphics."
    },
    {
      name: "minishell",
      pictures: [
        {
          img: require("./src/assets/minishell.png")
        }
      ],
      technologies: ["C", "Makefile", "Shell"],
      category: "Shell/Systems Programming",
      github: "https://github.com/Raainshe/minishell",
      date: "2025",
      visit: "",
      description:
        "A bash-like shell implementation for 42 School featuring command execution, pipes, redirections, environment variable expansion, signal handling, and built-in commands. Includes a custom tokenizer, a context-based garbage collector for memory management, and comprehensive tests for both functionality and memory leaks."
    }
  ],
  recommendations: [
    {
      title:
        "In his internship, Hrishikesh has demonstrated excellent learning ability, and with a dedicated, task oriented approach, he was able to complete his assignment ahead of time.",
      author: "Ushanas Shastri",
      position: "CTO",
      company: "Viteos Capital Market Services",
      location: "Mumbai"
    },
    {
      title:
        "I feel his budding leadership abilities will become even more effective in a diverse and challenging environment.",
      author: "Anil Dukkipatty",
      position: "CTO",
      company: "Elemential Labs",
      location: "Mumbai"
    },
    {
      title:
        "He has gained great knowledge and experience of SDE, and has learned ot develop the application keeping in mind the client's perpective and also making it creative.",
      author: "Chintan Shah",
      position: "Director",
      company: "Hridayam Soft Solution",
      location: "Mumbai"
    },
    {
      title:
        "During the course of his employment we have found him to be a self-started who was motivated, duty bound and a highly commited team player.",
      author: "Mrinal Pai",
      position: "Co-Founder & Director",
      company: "Skylark Drones",
      location: "Bangalore"
    }
  ]
};

export default info;
