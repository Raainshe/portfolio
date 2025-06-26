#!/bin/bash

# Script to update domain in meta tags for WhatsApp/Facebook preview
# Usage: ./update-domain.sh your-domain.com

if [ $# -eq 0 ]; then
    echo "Usage: $0 <your-domain.com>"
    echo "Example: $0 myportfolio.com"
    exit 1
fi

DOMAIN=$1

echo "Updating domain to: $DOMAIN"

# Update the index.html file
sed -i "s/YOUR-ACTUAL-DOMAIN\.com/$DOMAIN/g" public/index.html

echo "Domain updated successfully!"
echo "Make sure to rebuild your project after this change." 