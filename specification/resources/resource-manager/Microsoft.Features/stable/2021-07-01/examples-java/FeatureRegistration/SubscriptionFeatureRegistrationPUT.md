Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.resources.fluent.models.SubscriptionFeatureRegistrationInner;
import com.azure.resourcemanager.resources.models.SubscriptionFeatureRegistrationProperties;

/** Samples for SubscriptionFeatureRegistrations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/FeatureRegistration/SubscriptionFeatureRegistrationPUT.json
     */
    /**
     * Sample code: Creates a feature registration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsAFeatureRegistration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .featureClient()
            .getSubscriptionFeatureRegistrations()
            .createOrUpdateWithResponse(
                "subscriptionFeatureRegistrationGroupTestRG",
                "testFeature",
                new SubscriptionFeatureRegistrationInner()
                    .withProperties(new SubscriptionFeatureRegistrationProperties()),
                Context.NONE);
    }
}
```
