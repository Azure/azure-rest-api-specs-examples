Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for PartnerRegistrations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/PartnerRegistrations_CreateOrUpdate.json
     */
    /**
     * Sample code: PartnerRegistrations_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerRegistrationsCreateOrUpdate(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerRegistrations()
            .define("examplePartnerRegistrationName1")
            .withRegion("global")
            .withExistingResourceGroup("examplerg")
            .withTags(mapOf("key1", "value1", "key2", "Value2", "key3", "Value3"))
            .withPartnerName("ContosoCorp")
            .withPartnerResourceTypeName("ContosoCorp.Accounts")
            .withPartnerResourceTypeDisplayName("ContocoCorp Accounts DisplayName Text")
            .withPartnerResourceTypeDescription("ContocoCorp Accounts Description Text")
            .withSetupUri("https://www.example.com/setup.html")
            .withLogoUri("https://www.example.com/logo.png")
            .withAuthorizedAzureSubscriptionIds(Arrays.asList("d48566a8-2428-4a6c-8347-9675d09fb851"))
            .create();
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
```
