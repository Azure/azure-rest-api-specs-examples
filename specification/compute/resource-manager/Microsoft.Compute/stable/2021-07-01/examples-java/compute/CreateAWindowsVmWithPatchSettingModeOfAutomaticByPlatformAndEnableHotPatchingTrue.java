import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineInner;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.ManagedDiskParameters;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.PatchSettings;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.VirtualMachineSizeTypes;
import com.azure.resourcemanager.compute.models.WindowsConfiguration;
import com.azure.resourcemanager.compute.models.WindowsVMGuestPatchMode;
import java.util.Arrays;

/** Samples for VirtualMachines CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/CreateAWindowsVmWithPatchSettingModeOfAutomaticByPlatformAndEnableHotPatchingTrue.json
     */
    /**
     * Sample code: Create a Windows vm with a patch setting patchMode of AutomaticByPlatform and enableHotpatching set
     * to true.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAWindowsVmWithAPatchSettingPatchModeOfAutomaticByPlatformAndEnableHotpatchingSetToTrue(
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
                    .withHardwareProfile(new HardwareProfile().withVmSize(VirtualMachineSizeTypes.STANDARD_D1_V2))
                    .withStorageProfile(
                        new StorageProfile()
                            .withImageReference(
                                new ImageReference()
                                    .withPublisher("MicrosoftWindowsServer")
                                    .withOffer("WindowsServer")
                                    .withSku("2016-Datacenter")
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
                            .withAdminPassword("{your-password}")
                            .withWindowsConfiguration(
                                new WindowsConfiguration()
                                    .withProvisionVMAgent(true)
                                    .withEnableAutomaticUpdates(true)
                                    .withPatchSettings(
                                        new PatchSettings()
                                            .withPatchMode(WindowsVMGuestPatchMode.AUTOMATIC_BY_PLATFORM)
                                            .withEnableHotpatching(true))))
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
