import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.QueueServicePropertiesInner;
import com.azure.resourcemanager.storage.models.CorsRule;
import com.azure.resourcemanager.storage.models.CorsRuleAllowedMethodsItem;
import com.azure.resourcemanager.storage.models.CorsRules;
import java.util.Arrays;

/** Samples for QueueServices SetServiceProperties. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/QueueServicesPut.json
     */
    /**
     * Sample code: QueueServicesPut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queueServicesPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getQueueServices()
            .setServicePropertiesWithResponse(
                "res4410",
                "sto8607",
                new QueueServicePropertiesInner()
                    .withCors(
                        new CorsRules()
                            .withCorsRules(
                                Arrays
                                    .asList(
                                        new CorsRule()
                                            .withAllowedOrigins(
                                                Arrays.asList("http://www.contoso.com", "http://www.fabrikam.com"))
                                            .withAllowedMethods(
                                                Arrays
                                                    .asList(
                                                        CorsRuleAllowedMethodsItem.GET,
                                                        CorsRuleAllowedMethodsItem.HEAD,
                                                        CorsRuleAllowedMethodsItem.POST,
                                                        CorsRuleAllowedMethodsItem.OPTIONS,
                                                        CorsRuleAllowedMethodsItem.MERGE,
                                                        CorsRuleAllowedMethodsItem.PUT))
                                            .withMaxAgeInSeconds(100)
                                            .withExposedHeaders(Arrays.asList("x-ms-meta-*"))
                                            .withAllowedHeaders(
                                                Arrays.asList("x-ms-meta-abc", "x-ms-meta-data*", "x-ms-meta-target*")),
                                        new CorsRule()
                                            .withAllowedOrigins(Arrays.asList("*"))
                                            .withAllowedMethods(Arrays.asList(CorsRuleAllowedMethodsItem.GET))
                                            .withMaxAgeInSeconds(2)
                                            .withExposedHeaders(Arrays.asList("*"))
                                            .withAllowedHeaders(Arrays.asList("*")),
                                        new CorsRule()
                                            .withAllowedOrigins(
                                                Arrays.asList("http://www.abc23.com", "https://www.fabrikam.com/*"))
                                            .withAllowedMethods(
                                                Arrays
                                                    .asList(
                                                        CorsRuleAllowedMethodsItem.GET, CorsRuleAllowedMethodsItem.PUT))
                                            .withMaxAgeInSeconds(2000)
                                            .withExposedHeaders(
                                                Arrays.asList("x-ms-meta-abc", "x-ms-meta-data*", "x-ms-meta-target*"))
                                            .withAllowedHeaders(Arrays.asList("x-ms-meta-12345675754564*"))))),
                Context.NONE);
    }
}
