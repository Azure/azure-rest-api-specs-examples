
import com.azure.resourcemanager.storage.fluent.models.BlobServicePropertiesInner;
import com.azure.resourcemanager.storage.models.ChangeFeed;
import com.azure.resourcemanager.storage.models.CorsRule;
import com.azure.resourcemanager.storage.models.CorsRuleAllowedMethodsItem;
import com.azure.resourcemanager.storage.models.CorsRules;
import com.azure.resourcemanager.storage.models.DeleteRetentionPolicy;
import com.azure.resourcemanager.storage.models.StaticWebsite;
import java.util.Arrays;

/**
 * Samples for BlobServices SetServiceProperties.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/BlobServicesPut.json
     */
    /**
     * Sample code: PutBlobServices.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void putBlobServices(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobServices().setServicePropertiesWithResponse("res4410", "sto8607",
            new BlobServicePropertiesInner().withCors(new CorsRules().withCorsRules(Arrays.asList(
                new CorsRule().withAllowedOrigins(Arrays.asList("http://www.contoso.com", "http://www.fabrikam.com"))
                    .withAllowedMethods(Arrays.asList(CorsRuleAllowedMethodsItem.GET, CorsRuleAllowedMethodsItem.HEAD,
                        CorsRuleAllowedMethodsItem.POST, CorsRuleAllowedMethodsItem.OPTIONS,
                        CorsRuleAllowedMethodsItem.MERGE, CorsRuleAllowedMethodsItem.PUT))
                    .withMaxAgeInSeconds(100).withExposedHeaders(Arrays.asList("x-ms-meta-*"))
                    .withAllowedHeaders(Arrays.asList("x-ms-meta-abc", "x-ms-meta-data*", "x-ms-meta-target*")),
                new CorsRule().withAllowedOrigins(Arrays.asList("*"))
                    .withAllowedMethods(Arrays.asList(CorsRuleAllowedMethodsItem.GET)).withMaxAgeInSeconds(2)
                    .withExposedHeaders(Arrays.asList("*")).withAllowedHeaders(Arrays.asList("*")),
                new CorsRule().withAllowedOrigins(Arrays.asList("http://www.abc23.com", "https://www.fabrikam.com/*"))
                    .withAllowedMethods(Arrays.asList(CorsRuleAllowedMethodsItem.GET, CorsRuleAllowedMethodsItem.PUT))
                    .withMaxAgeInSeconds(2000)
                    .withExposedHeaders(Arrays.asList("x-ms-meta-abc", "x-ms-meta-data*", "x -ms-meta-target*"))
                    .withAllowedHeaders(Arrays.asList("x-ms-meta-12345675754564*")))))
                .withDefaultServiceVersion("2017-07-29")
                .withDeleteRetentionPolicy(new DeleteRetentionPolicy().withEnabled(true).withDays(300))
                .withStaticWebsite(new StaticWebsite().withEnabled(true).withIndexDocument("home.html")
                    .withErrorDocument404Path("site/errors/not-found.html"))
                .withIsVersioningEnabled(true).withChangeFeed(
                    new ChangeFeed().withEnabled(true).withRetentionInDays(7)),
            com.azure.core.util.Context.NONE);
    }
}
