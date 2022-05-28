Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.OrchestrationServiceNames;
import com.azure.resourcemanager.compute.models.OrchestrationServiceStateAction;
import com.azure.resourcemanager.compute.models.OrchestrationServiceStateInput;

/** Samples for VirtualMachineScaleSets SetOrchestrationServiceState. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_SetOrchestrationServiceState_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_SetOrchestrationServiceState_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsSetOrchestrationServiceStateMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .setOrchestrationServiceState(
                "rgcompute",
                "aaaaaaaaaaaaaaaaaaaaaaaa",
                new OrchestrationServiceStateInput()
                    .withServiceName(OrchestrationServiceNames.AUTOMATIC_REPAIRS)
                    .withAction(OrchestrationServiceStateAction.RESUME),
                Context.NONE);
    }
}
```
