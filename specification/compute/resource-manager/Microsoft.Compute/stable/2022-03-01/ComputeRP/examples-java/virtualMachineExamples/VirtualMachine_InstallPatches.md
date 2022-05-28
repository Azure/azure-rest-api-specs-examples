Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.VMGuestPatchClassificationWindows;
import com.azure.resourcemanager.compute.models.VMGuestPatchRebootSetting;
import com.azure.resourcemanager.compute.models.VirtualMachineInstallPatchesParameters;
import com.azure.resourcemanager.compute.models.WindowsParameters;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for VirtualMachines InstallPatches. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachine_InstallPatches.json
     */
    /**
     * Sample code: Install patch state of a virtual machine.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void installPatchStateOfAVirtualMachine(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .installPatches(
                "myResourceGroupName",
                "myVMName",
                new VirtualMachineInstallPatchesParameters()
                    .withMaximumDuration("PT4H")
                    .withRebootSetting(VMGuestPatchRebootSetting.IF_REQUIRED)
                    .withWindowsParameters(
                        new WindowsParameters()
                            .withClassificationsToInclude(
                                Arrays
                                    .asList(
                                        VMGuestPatchClassificationWindows.CRITICAL,
                                        VMGuestPatchClassificationWindows.SECURITY))
                            .withMaxPatchPublishDate(OffsetDateTime.parse("2020-11-19T02:36:43.0539904+00:00"))),
                Context.NONE);
    }
}
```
