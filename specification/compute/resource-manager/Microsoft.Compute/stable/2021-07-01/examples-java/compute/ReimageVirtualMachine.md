```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.VirtualMachineReimageParameters;

/** Samples for VirtualMachines Reimage. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/ReimageVirtualMachine.json
     */
    /**
     * Sample code: Reimage a Virtual Machine.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reimageAVirtualMachine(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .reimage(
                "myResourceGroup", "myVMName", new VirtualMachineReimageParameters().withTempDisk(true), Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
