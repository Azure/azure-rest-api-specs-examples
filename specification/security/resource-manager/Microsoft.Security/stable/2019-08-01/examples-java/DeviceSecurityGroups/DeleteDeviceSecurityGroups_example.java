
/**
 * Samples for DeviceSecurityGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/
     * DeleteDeviceSecurityGroups_example.json
     */
    /**
     * Sample code: Delete a device security group for the specified IoT Hub resource.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteADeviceSecurityGroupForTheSpecifiedIoTHubResource(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.deviceSecurityGroups().deleteByResourceGroupWithResponse(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Devices/iotHubs/sampleiothub",
            "samplesecuritygroup", com.azure.core.util.Context.NONE);
    }
}
