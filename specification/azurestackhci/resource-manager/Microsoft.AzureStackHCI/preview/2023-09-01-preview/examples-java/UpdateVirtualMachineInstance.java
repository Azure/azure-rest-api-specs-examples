import com.azure.resourcemanager.azurestackhci.models.StorageProfileUpdate;
import com.azure.resourcemanager.azurestackhci.models.StorageProfileUpdateDataDisksItem;
import com.azure.resourcemanager.azurestackhci.models.VirtualMachineInstanceUpdateProperties;
import com.azure.resourcemanager.azurestackhci.models.VirtualMachineInstanceUpdateRequest;
import java.util.Arrays;

/** Samples for VirtualMachineInstances Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/UpdateVirtualMachineInstance.json
     */
    /**
     * Sample code: UpdateVirtualMachine.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void updateVirtualMachine(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .virtualMachineInstances()
            .update(
                "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM",
                new VirtualMachineInstanceUpdateRequest()
                    .withProperties(
                        new VirtualMachineInstanceUpdateProperties()
                            .withStorageProfile(
                                new StorageProfileUpdate()
                                    .withDataDisks(
                                        Arrays
                                            .asList(
                                                new StorageProfileUpdateDataDisksItem()
                                                    .withId(
                                                        "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd"))))),
                com.azure.core.util.Context.NONE);
    }
}
