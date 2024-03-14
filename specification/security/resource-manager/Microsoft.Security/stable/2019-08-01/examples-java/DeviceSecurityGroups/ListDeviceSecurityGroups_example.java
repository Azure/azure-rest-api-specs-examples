
/**
 * Samples for DeviceSecurityGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/
     * ListDeviceSecurityGroups_example.json
     */
    /**
     * Sample code: List all device security groups for the specified IoT Hub resource.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listAllDeviceSecurityGroupsForTheSpecifiedIoTHubResource(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.deviceSecurityGroups().list(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Devices/iotHubs/sampleiothub",
            com.azure.core.util.Context.NONE);
    }
}
