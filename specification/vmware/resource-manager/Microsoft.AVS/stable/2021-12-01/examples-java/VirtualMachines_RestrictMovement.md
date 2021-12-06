Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.avs.models.VirtualMachineRestrictMovement;
import com.azure.resourcemanager.avs.models.VirtualMachineRestrictMovementState;

/** Samples for VirtualMachines RestrictMovement. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/VirtualMachines_RestrictMovement.json
     */
    /**
     * Sample code: VirtualMachine_RestrictMovement.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void virtualMachineRestrictMovement(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .virtualMachines()
            .restrictMovement(
                "group1",
                "cloud1",
                "cluster1",
                "vm-209",
                new VirtualMachineRestrictMovement().withRestrictMovement(VirtualMachineRestrictMovementState.ENABLED),
                Context.NONE);
    }
}
```
