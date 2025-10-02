<?xml version="1.0" encoding="utf-8"?>
<xsl:stylesheet version="3.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
    <xsl:output method="html" version="1.0" encoding="UTF-8" indent="yes"/>
    <xsl:template match="/">
        <html xmlns="http://www.w3.org/1999/xhtml" lang="en">
            <head>
                <title>
                    <xsl:value-of select="/rss/channel/title"/>
                </title>
                <meta charset="utf-8"/>
                <meta name="viewport" content="width=device-width, initial-scale=1"/>
                <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.8/dist/css/bootstrap.min.css" rel="stylesheet"
                      integrity="sha384-sRIl4kxILFvY47J16cr9ZwB07vP4J8+LH7qKQnuqkuIAvNWLzeN8tE5YBujZqJLB"
                      crossorigin="anonymous"/>
            </head>
            <body class="bg-white">
                <main class="container-md px-3 py-3 markdown-body">
                    <header class="py-5">
                        <h1 class="border-0">
                            <xsl:value-of select="/rss/channel/title"/>
                        </h1>
                        <p>
                            <xsl:value-of select="/rss/channel/description"/>
                        </p>
                        <a class="text-decoration-none" target="_blank">
                            <xsl:attribute name="href">
                                <xsl:value-of select="/rss/channel/link"/>
                            </xsl:attribute>
                            Visit Website &#x2192;
                        </a>
                    </header>
                    <h2>Recent Items</h2>
                    <xsl:for-each select="/rss/channel/item">
                        <section class="pb-5">
                            <h3 class="mb-0">
                                <a class="text-decoration-none" target="_blank">
                                    <xsl:attribute name="href">
                                        <xsl:value-of select="link"/>
                                    </xsl:attribute>
                                    <xsl:value-of select="title"/>
                                </a>
                            </h3>
                            <p>
                                Published:
                                <xsl:value-of select="pubDate"/>
                            </p>
                        </section>
                    </xsl:for-each>
                </main>
            </body>
        </html>
    </xsl:template>
</xsl:stylesheet>