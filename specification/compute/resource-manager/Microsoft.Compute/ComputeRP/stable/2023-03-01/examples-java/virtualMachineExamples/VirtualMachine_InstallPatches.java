import com.azure.resourcemanager.compute.models.VMGuestPatchClassificationWindows;
import com.azure.resourcemanager.compute.models.VMGuestPatchRebootSetting;
import com.azure.resourcemanager.compute.models.VirtualMachineInstallPatchesParameters;
import com.azure.resourcemanager.compute.models.WindowsParameters;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for VirtualMachines InstallPatches. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineExamples/VirtualMachine_InstallPatches.json
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
                com.azure.core.util.Context.NONE);
    }
}
