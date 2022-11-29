import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineInner;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.LinuxConfiguration;
import com.azure.resourcemanager.compute.models.LinuxPatchAssessmentMode;
import com.azure.resourcemanager.compute.models.LinuxPatchSettings;
import com.azure.resourcemanager.compute.models.LinuxVMGuestPatchAutomaticByPlatformRebootSetting;
import com.azure.resourcemanager.compute.models.LinuxVMGuestPatchAutomaticByPlatformSettings;
import com.azure.resourcemanager.compute.models.LinuxVMGuestPatchMode;
import com.azure.resourcemanager.compute.models.ManagedDiskParameters;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineSizeTypes;
import java.util.Arrays;

/** Samples for VirtualMachines CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachine_Create_LinuxVmWithAutomaticByPlatformSettings.json
     */
    /**
     * Sample code: Create a Linux vm with a patch setting patchMode of AutomaticByPlatform and
     * AutomaticByPlatformSettings.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createALinuxVmWithAPatchSettingPatchModeOfAutomaticByPlatformAndAutomaticByPlatformSettings(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .createOrUpdate(
                "myResourceGroup",
                "myVM",
                new VirtualMachineInner()
                    .withLocation("westus")
                    .withHardwareProfile(new HardwareProfile().withVmSize(VirtualMachineSizeTypes.STANDARD_D2S_V3))
                    .withStorageProfile(
                        new StorageProfile()
                            .withImageReference(
                                new ImageReference()
                                    .withPublisher("Canonical")
                                    .withOffer("UbuntuServer")
                                    .withSku("16.04-LTS")
                                    .withVersion("latest"))
                            .withOsDisk(
                                new OSDisk()
                                    .withName("myVMosdisk")
                                    .withCaching(CachingTypes.READ_WRITE)
                                    .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE)
                                    .withManagedDisk(
                                        new ManagedDiskParameters()
                                            .withStorageAccountType(StorageAccountTypes.PREMIUM_LRS))))
                    .withOsProfile(
                        new OSProfile()
                            .withComputerName("myVM")
                            .withAdminUsername("{your-username}")
                            .withAdminPassword("fakeTokenPlaceholder")
                            .withLinuxConfiguration(
                                new LinuxConfiguration()
                                    .withProvisionVMAgent(true)
                                    .withPatchSettings(
                                        new LinuxPatchSettings()
                                            .withPatchMode(LinuxVMGuestPatchMode.AUTOMATIC_BY_PLATFORM)
                                            .withAssessmentMode(LinuxPatchAssessmentMode.AUTOMATIC_BY_PLATFORM)
                                            .withAutomaticByPlatformSettings(
                                                new LinuxVMGuestPatchAutomaticByPlatformSettings()
                                                    .withRebootSetting(
                                                        LinuxVMGuestPatchAutomaticByPlatformRebootSetting.NEVER)))))
                    .withNetworkProfile(
                        new NetworkProfile()
                            .withNetworkInterfaces(
                                Arrays
                                    .asList(
                                        new NetworkInterfaceReference()
                                            .withId(
                                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}")
                                            .withPrimary(true)))),
                Context.NONE);
    }
}
