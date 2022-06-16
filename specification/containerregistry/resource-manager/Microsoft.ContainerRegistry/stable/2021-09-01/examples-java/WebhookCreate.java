import com.azure.core.util.Context;
import com.azure.resourcemanager.containerregistry.models.WebhookAction;
import com.azure.resourcemanager.containerregistry.models.WebhookCreateParameters;
import com.azure.resourcemanager.containerregistry.models.WebhookStatus;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Webhooks Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/WebhookCreate.json
     */
    /**
     * Sample code: WebhookCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void webhookCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getWebhooks()
            .create(
                "myResourceGroup",
                "myRegistry",
                "myWebhook",
                new WebhookCreateParameters()
                    .withTags(mapOf("key", "value"))
                    .withLocation("westus")
                    .withServiceUri("http://myservice.com")
                    .withCustomHeaders(
                        mapOf("Authorization", "Basic 000000000000000000000000000000000000000000000000000"))
                    .withStatus(WebhookStatus.ENABLED)
                    .withScope("myRepository")
                    .withActions(Arrays.asList(WebhookAction.PUSH)),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
