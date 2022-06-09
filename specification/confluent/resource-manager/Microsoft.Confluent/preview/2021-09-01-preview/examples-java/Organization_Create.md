```java
import com.azure.resourcemanager.confluent.models.OfferDetail;
import com.azure.resourcemanager.confluent.models.UserDetail;
import java.util.HashMap;
import java.util.Map;

/** Samples for Organization Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_Create.json
     */
    /**
     * Sample code: Organization_Create.
     *
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationCreate(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager
            .organizations()
            .define("myOrganization")
            .withRegion("West US")
            .withExistingResourceGroup("myResourceGroup")
            .withOfferDetail(
                new OfferDetail()
                    .withPublisherId("string")
                    .withId("string")
                    .withPlanId("string")
                    .withPlanName("string")
                    .withTermUnit("string"))
            .withUserDetail(
                new UserDetail()
                    .withFirstName("string")
                    .withLastName("string")
                    .withEmailAddress("contoso@microsoft.com"))
            .withTags(mapOf("Environment", "Dev"))
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-confluent_1.0.0-beta.3/sdk/confluent/azure-resourcemanager-confluent/README.md) on how to add the SDK to your project and authenticate.
