import com.azure.resourcemanager.azurestackhci.fluent.models.VirtualMachineInstanceInner;
import com.azure.resourcemanager.azurestackhci.models.ExtendedLocation;
import com.azure.resourcemanager.azurestackhci.models.ExtendedLocationTypes;
import com.azure.resourcemanager.azurestackhci.models.VirtualMachineInstancePropertiesHardwareProfile;
import com.azure.resourcemanager.azurestackhci.models.VirtualMachineInstancePropertiesNetworkProfile;
import com.azure.resourcemanager.azurestackhci.models.VirtualMachineInstancePropertiesNetworkProfileNetworkInterfacesItem;
import com.azure.resourcemanager.azurestackhci.models.VirtualMachineInstancePropertiesSecurityProfile;
import com.azure.resourcemanager.azurestackhci.models.VirtualMachineInstancePropertiesSecurityProfileUefiSettings;
import com.azure.resourcemanager.azurestackhci.models.VirtualMachineInstancePropertiesStorageProfile;
import com.azure.resourcemanager.azurestackhci.models.VirtualMachineInstancePropertiesStorageProfileOsDisk;
import com.azure.resourcemanager.azurestackhci.models.VmSizeEnum;
import java.util.Arrays;

/** Samples for VirtualMachineInstances CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutVirtualMachineInstanceWithOsDisk.json
     */
    /**
     * Sample code: PutVirtualMachineInstanceWithOsDisk.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void putVirtualMachineInstanceWithOsDisk(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .virtualMachineInstances()
            .createOrUpdate(
                "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM",
                new VirtualMachineInstanceInner()
                    .withExtendedLocation(
                        new ExtendedLocation()
                            .withName(
                                "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location")
                            .withType(ExtendedLocationTypes.CUSTOM_LOCATION))
                    .withHardwareProfile(
                        new VirtualMachineInstancePropertiesHardwareProfile().withVmSize(VmSizeEnum.DEFAULT))
                    .withNetworkProfile(
                        new VirtualMachineInstancePropertiesNetworkProfile()
                            .withNetworkInterfaces(
                                Arrays
                                    .asList(
                                        new VirtualMachineInstancePropertiesNetworkProfileNetworkInterfacesItem()
                                            .withId("test-nic"))))
                    .withSecurityProfile(
                        new VirtualMachineInstancePropertiesSecurityProfile()
                            .withEnableTpm(true)
                            .withUefiSettings(
                                new VirtualMachineInstancePropertiesSecurityProfileUefiSettings()
                                    .withSecureBootEnabled(true)))
                    .withStorageProfile(
                        new VirtualMachineInstancePropertiesStorageProfile()
                            .withOsDisk(
                                new VirtualMachineInstancePropertiesStorageProfileOsDisk()
                                    .withId(
                                        "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd"))
                            .withVmConfigStoragePathId(
                                "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-container")),
                com.azure.core.util.Context.NONE);
    }
}
