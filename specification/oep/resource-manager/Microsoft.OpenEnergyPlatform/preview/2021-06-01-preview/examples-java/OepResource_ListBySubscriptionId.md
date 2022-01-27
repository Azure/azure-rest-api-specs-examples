Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-oep_1.0.0-beta.1/sdk/oep/azure-resourcemanager-oep/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for EnergyServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/OepResource_ListBySubscriptionId.json
     */
    /**
     * Sample code: OepResource_ListBySubscriptionId.
     *
     * @param manager Entry point to OepManager.
     */
    public static void oepResourceListBySubscriptionId(com.azure.resourcemanager.oep.OepManager manager) {
        manager.energyServices().list(Context.NONE);
    }
}
```
