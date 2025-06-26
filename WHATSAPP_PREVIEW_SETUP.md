# WhatsApp Link Preview Setup

This document explains how to set up proper link previews for WhatsApp, Facebook, Twitter, and other social media platforms.

## What's Been Added

The following meta tags have been added to `public/index.html`:

### Open Graph Tags (for WhatsApp, Facebook, LinkedIn)
- `og:title` - Portfolio title
- `og:description` - Professional description
- `og:image` - Profile image (ryan.jpg)
- `og:url` - Website URL
- `og:type` - Website type
- `og:image:width` and `og:image:height` - Image dimensions
- `og:image:alt` - Alt text for accessibility

### Twitter Card Tags
- `twitter:card` - Large image card type
- `twitter:title` - Portfolio title
- `twitter:description` - Professional description
- `twitter:image` - Profile image
- `twitter:url` - Website URL

### SEO Meta Tags
- Title, description, keywords
- Author information
- Robots directives
- Language settings

## Setup Instructions

### 1. Update Your Domain
Before deploying, update the domain in the meta tags:

```bash
./update-domain.sh your-actual-domain.com
```

Replace `your-actual-domain.com` with your actual domain (e.g., `ryanmakoni.com`).

### 2. Build and Deploy
After updating the domain:

```bash
npm run build
```

Deploy the contents of the `dist` folder to your hosting provider.

### 3. Test the Preview
You can test your link preview using these tools:

- **Facebook Sharing Debugger**: https://developers.facebook.com/tools/debug/
- **Twitter Card Validator**: https://cards-dev.twitter.com/validator
- **LinkedIn Post Inspector**: https://www.linkedin.com/post-inspector/
- **WhatsApp**: Share the link in a WhatsApp chat to see the preview

## Image Requirements

The profile image (`ryan.jpg`) should meet these requirements for optimal display:

- **Recommended size**: 1200x630 pixels (1.91:1 aspect ratio)
- **Minimum size**: 600x315 pixels
- **Maximum file size**: 8MB
- **Format**: JPG, PNG, or GIF

## Troubleshooting

### Preview Not Showing
1. Make sure your domain is correctly set in the meta tags
2. Clear the cache using Facebook's Sharing Debugger
3. Ensure the image URL is accessible (try opening it directly in a browser)
4. Check that the image meets the size requirements

### Wrong Image Showing
1. Clear the cache using the social media platform's debugger
2. Ensure the image URL is correct and accessible
3. Wait 24-48 hours for cache to refresh (some platforms cache aggressively)

### Wrong Title/Description
1. Update the meta tags in `public/index.html`
2. Rebuild and redeploy the site
3. Clear the cache using the platform's debugger

## Maintenance

- Update the meta description periodically to reflect current skills/projects
- Keep the profile image current and professional
- Test link previews after major updates
- Monitor social media analytics to see how your links are being shared

## Files Modified

- `public/index.html` - Added comprehensive meta tags
- `public/ryan.jpg` - Profile image for previews
- `update-domain.sh` - Script to easily update domain
- `WHATSAPP_PREVIEW_SETUP.md` - This documentation file 