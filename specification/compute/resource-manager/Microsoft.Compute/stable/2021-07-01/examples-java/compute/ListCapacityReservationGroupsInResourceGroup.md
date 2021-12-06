Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.ExpandTypesForGetCapacityReservationGroups;

/** Samples for CapacityReservationGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/ListCapacityReservationGroupsInResourceGroup.json
     */
    /**
     * Sample code: List capacity reservation groups in resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCapacityReservationGroupsInResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCapacityReservationGroups()
            .listByResourceGroup(
                "myResourceGroup", ExpandTypesForGetCapacityReservationGroups.VIRTUAL_MACHINES_REF, Context.NONE);
    }
}
```
