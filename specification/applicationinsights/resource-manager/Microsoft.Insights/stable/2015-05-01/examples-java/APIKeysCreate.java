import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.ApiKeyRequest;
import java.util.Arrays;

/** Samples for ApiKeys Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/APIKeysCreate.json
     */
    /**
     * Sample code: APIKeyCreate.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void aPIKeyCreate(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .apiKeys()
            .createWithResponse(
                "my-resource-group",
                "my-component",
                new ApiKeyRequest()
                    .withName("test2")
                    .withLinkedReadProperties(
                        Arrays
                            .asList(
                                "/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/api",
                                "/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/agentconfig"))
                    .withLinkedWriteProperties(
                        Arrays
                            .asList(
                                "/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/annotations")),
                Context.NONE);
    }
}
