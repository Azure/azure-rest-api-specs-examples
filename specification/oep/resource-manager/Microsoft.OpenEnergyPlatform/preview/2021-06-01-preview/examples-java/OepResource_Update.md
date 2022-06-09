```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.oep.models.EnergyService;

/** Samples for EnergyServices Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/oep/resource-manager/Microsoft.OpenEnergyPlatform/preview/2021-06-01-preview/examples/OepResource_Update.json
     */
    /**
     * Sample code: OepResource_Update.
     *
     * @param manager Entry point to OepManager.
     */
    public static void oepResourceUpdate(com.azure.resourcemanager.oep.OepManager manager) {
        EnergyService resource =
            manager
                .energyServices()
                .getByResourceGroupWithResponse("DummyResourceGroupName", "DummyResourceName", Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-oep_1.0.0-beta.1/sdk/oep/azure-resourcemanager-oep/README.md) on how to add the SDK to your project and authenticate.
