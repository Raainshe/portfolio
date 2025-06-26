# SSL Certificate & 403 Error Troubleshooting Guide

## Common Causes of 403 Forbidden After SSL Setup

### 1. **Nginx Configuration Issues**
The most common cause is that nginx isn't properly configured for SSL.

**Solution**: Use the updated `nginx_default.conf` file I provided.

### 2. **Certificate Path Issues**
Make sure your SSL certificate paths are correct.

**Check certificate location**:
```bash
sudo find /etc -name "*.pem" | grep ryanmakoni
```

**Common certificate paths**:
- Let's Encrypt: `/etc/letsencrypt/live/ryanmakoni.live/`
- Certbot: `/etc/letsencrypt/live/ryanmakoni.live/`
- Manual: Check your certificate provider's documentation

### 3. **File Permissions**
SSL certificates need proper permissions.

**Fix permissions**:
```bash
sudo chmod 644 /etc/letsencrypt/live/ryanmakoni.live/fullchain.pem
sudo chmod 600 /etc/letsencrypt/live/ryanmakoni.live/privkey.pem
sudo chown root:root /etc/letsencrypt/live/ryanmakoni.live/fullchain.pem
sudo chown root:root /etc/letsencrypt/live/ryanmakoni.live/privkey.pem
```

### 4. **Nginx Service Issues**
Restart nginx after configuration changes.

```bash
sudo nginx -t  # Test configuration
sudo systemctl restart nginx
sudo systemctl status nginx  # Check if running
```

## Step-by-Step Fix

### Step 1: Update Nginx Configuration
1. Copy the updated `nginx_default.conf` to your server
2. Replace your current nginx configuration with this file
3. Update the SSL certificate paths if they're different

### Step 2: Check Certificate Paths
```bash
# Find your certificate files
sudo find /etc -name "fullchain.pem" 2>/dev/null
sudo find /etc -name "privkey.pem" 2>/dev/null
```

### Step 3: Test and Restart Nginx
```bash
sudo nginx -t
sudo systemctl restart nginx
```

### Step 4: Check Logs
If still having issues, check the logs:
```bash
sudo tail -f /var/log/nginx/error.log
sudo tail -f /var/log/nginx/access.log
```

## Alternative Solutions

### If Using Certbot
If you used Certbot to generate certificates, it might have created its own configuration:

```bash
# Check if certbot created a configuration
sudo ls /etc/nginx/sites-available/
sudo ls /etc/nginx/sites-enabled/
```

### If Using Different Certificate Provider
Update the SSL certificate paths in the nginx configuration:

```nginx
ssl_certificate /path/to/your/certificate.crt;
ssl_certificate_key /path/to/your/private.key;
```

### If Using Cloudflare or CDN
If you're using Cloudflare or another CDN:
1. Make sure SSL/TLS encryption mode is set to "Full" or "Full (strict)"
2. Check that the origin certificate is properly configured

## Testing Your Setup

### Test SSL Configuration
```bash
# Test from command line
curl -I https://ryanmakoni.live

# Test SSL certificate
openssl s_client -connect ryanmakoni.live:443 -servername ryanmakoni.live
```

### Test HTTP to HTTPS Redirect
```bash
curl -I http://ryanmakoni.live
# Should return 301 redirect to HTTPS
```

## Common Error Messages

### "SSL certificate not found"
- Check certificate file paths
- Verify file permissions
- Ensure nginx can read the certificate files

### "Permission denied"
- Fix file permissions on certificate files
- Check nginx user permissions

### "No such file or directory"
- Verify certificate file paths
- Check if certificates were properly installed

## Emergency Fallback

If you need to temporarily disable SSL to get the site working:

```nginx
server {
    listen 80;
    listen [::]:80;
    server_name ryanmakoni.live www.ryanmakoni.live;
    root /var/www/html;
    index index.html index.htm;
    location / {
        try_files $uri $uri/ /index.html;
    }
}
```

## Getting Help

If you're still having issues:
1. Check nginx error logs: `sudo tail -f /var/log/nginx/error.log`
2. Verify certificate installation: `sudo certbot certificates`
3. Test nginx configuration: `sudo nginx -t`
4. Check system resources: `sudo systemctl status nginx` 