import info_en from '../../info.js';
import info_de from '../../info_de.js';

class LanguageService {
  constructor() {
    this.currentLanguage = 'en';
    this.languages = {
      en: info_en,
      de: info_de
    };
    this.translations = {
      en: {
        navbar: {
          about: 'about',
          skills: 'skills',
          portfolio: 'portfolio'
        },
        home: {
          greeting: 'hello there!'
        },
        about: {
          title: 'about me.',
          education: 'education',
          experiences: 'experiences'
        },
        skills: {
          title: 'skills.'
        },
        portfolio: {
          title: 'portfolio.',
          development: 'development',
          personal_projects: 'personal projects/42 projects',
          read_more: 'read more'
        }
      },
      de: {
        navbar: {
          about: 'über mich',
          skills: 'fähigkeiten',
          portfolio: 'portfolio'
        },
        home: {
          greeting: 'hallo!'
        },
        about: {
          title: 'über mich.',
          education: 'bildung',
          experiences: 'erfahrungen'
        },
        skills: {
          title: 'fähigkeiten.'
        },
        portfolio: {
          title: 'portfolio.',
          development: 'entwicklung',
          personal_projects: 'persönliche projekte/42 projekte',
          read_more: 'mehr lesen'
        }
      }
    };
    this.listeners = [];
  }

  getCurrentLanguage() {
    return this.currentLanguage;
  }

  setLanguage(language) {
    if (this.languages[language]) {
      this.currentLanguage = language;
      this.notifyListeners();
      // Store in localStorage for persistence
      localStorage.setItem('preferred-language', language);
    }
  }

  getInfo() {
    return this.languages[this.currentLanguage];
  }

  getTranslations() {
    return this.translations[this.currentLanguage];
  }

  subscribe(callback) {
    this.listeners.push(callback);
  }

  unsubscribe(callback) {
    this.listeners = this.listeners.filter(listener => listener !== callback);
  }

  notifyListeners() {
    this.listeners.forEach(callback => callback(this.currentLanguage, this.getInfo()));
  }

  // Initialize language from localStorage or browser preference
  init() {
    const savedLanguage = localStorage.getItem('preferred-language');
    const browserLanguage = navigator.language.split('-')[0];
    
    if (savedLanguage && this.languages[savedLanguage]) {
      this.setLanguage(savedLanguage);
    } else if (this.languages[browserLanguage]) {
      this.setLanguage(browserLanguage);
    } else {
      this.setLanguage('en'); // Default fallback
    }
  }
}

// Create a singleton instance
const languageService = new LanguageService();
export default languageService;
