
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineInner;
import com.azure.resourcemanager.compute.models.BootDiagnostics;
import com.azure.resourcemanager.compute.models.CachingTypes;
import com.azure.resourcemanager.compute.models.DiagnosticsProfile;
import com.azure.resourcemanager.compute.models.DiskCreateOptionTypes;
import com.azure.resourcemanager.compute.models.EventGridAndResourceGraph;
import com.azure.resourcemanager.compute.models.HardwareProfile;
import com.azure.resourcemanager.compute.models.ImageReference;
import com.azure.resourcemanager.compute.models.ManagedDiskParameters;
import com.azure.resourcemanager.compute.models.NetworkInterfaceReference;
import com.azure.resourcemanager.compute.models.NetworkProfile;
import com.azure.resourcemanager.compute.models.OSDisk;
import com.azure.resourcemanager.compute.models.OSImageNotificationProfile;
import com.azure.resourcemanager.compute.models.OSProfile;
import com.azure.resourcemanager.compute.models.ScheduledEventsAdditionalPublishingTargets;
import com.azure.resourcemanager.compute.models.ScheduledEventsPolicy;
import com.azure.resourcemanager.compute.models.ScheduledEventsProfile;
import com.azure.resourcemanager.compute.models.StorageAccountTypes;
import com.azure.resourcemanager.compute.models.StorageProfile;
import com.azure.resourcemanager.compute.models.TerminateNotificationProfile;
import com.azure.resourcemanager.compute.models.UserInitiatedReboot;
import com.azure.resourcemanager.compute.models.UserInitiatedRedeploy;
import com.azure.resourcemanager.compute.models.VirtualMachineSizeTypes;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualMachines CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineExamples/VirtualMachine_Create_WithScheduledEventsProfile.json
     */
    /**
     * Sample code: Create a vm with Scheduled Events Profile.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAVmWithScheduledEventsProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().createOrUpdate("myResourceGroup", "myVM",
            new VirtualMachineInner().withLocation("westus")
                .withHardwareProfile(new HardwareProfile().withVmSize(VirtualMachineSizeTypes.STANDARD_D1_V2))
                .withScheduledEventsPolicy(new ScheduledEventsPolicy()
                    .withUserInitiatedRedeploy(new UserInitiatedRedeploy().withAutomaticallyApprove(true))
                    .withUserInitiatedReboot(new UserInitiatedReboot().withAutomaticallyApprove(true))
                    .withScheduledEventsAdditionalPublishingTargets(new ScheduledEventsAdditionalPublishingTargets()
                        .withEventGridAndResourceGraph(new EventGridAndResourceGraph().withEnable(true))))
                .withStorageProfile(new StorageProfile()
                    .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                        .withOffer("WindowsServer").withSku("2016-Datacenter").withVersion("latest"))
                    .withOsDisk(new OSDisk().withName("myVMosdisk").withCaching(CachingTypes.READ_WRITE)
                        .withCreateOption(DiskCreateOptionTypes.FROM_IMAGE).withManagedDisk(
                            new ManagedDiskParameters().withStorageAccountType(StorageAccountTypes.STANDARD_LRS))))
                .withOsProfile(new OSProfile().withComputerName("myVM").withAdminUsername("{your-username}")
                    .withAdminPassword("fakeTokenPlaceholder"))
                .withNetworkProfile(
                    new NetworkProfile().withNetworkInterfaces(Arrays.asList(new NetworkInterfaceReference().withId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}")
                        .withPrimary(true))))
                .withDiagnosticsProfile(new DiagnosticsProfile().withBootDiagnostics(new BootDiagnostics()
                    .withEnabled(true).withStorageUri("http://{existing-storage-account-name}.blob.core.windows.net")))
                .withScheduledEventsProfile(new ScheduledEventsProfile()
                    .withTerminateNotificationProfile(
                        new TerminateNotificationProfile().withNotBeforeTimeout("PT10M").withEnable(true))
                    .withOsImageNotificationProfile(
                        new OSImageNotificationProfile().withNotBeforeTimeout("PT15M").withEnable(true))),
            null, null, com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
