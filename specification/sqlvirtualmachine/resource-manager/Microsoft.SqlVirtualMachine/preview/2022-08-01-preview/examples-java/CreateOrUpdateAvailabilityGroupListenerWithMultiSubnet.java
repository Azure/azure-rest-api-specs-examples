
import com.azure.resourcemanager.sqlvirtualmachine.models.MultiSubnetIpConfiguration;
import com.azure.resourcemanager.sqlvirtualmachine.models.PrivateIpAddress;
import java.util.Arrays;

/**
 * Samples for AvailabilityGroupListeners CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/
     * CreateOrUpdateAvailabilityGroupListenerWithMultiSubnet.json
     */
    /**
     * Sample code: Creates or updates an availability group listener. This is used for VMs present in multi subnet.
     * 
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void createsOrUpdatesAnAvailabilityGroupListenerThisIsUsedForVMsPresentInMultiSubnet(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.availabilityGroupListeners().define("agl-test")
            .withExistingSqlVirtualMachineGroup("testrg", "testvmgroup").withAvailabilityGroupName("ag-test")
            .withMultiSubnetIpConfigurations(Arrays.asList(new MultiSubnetIpConfiguration().withPrivateIpAddress(
                new PrivateIpAddress().withIpAddress("10.0.0.112").withSubnetResourceId(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"))
                .withSqlVirtualMachineInstance(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2"),
                new MultiSubnetIpConfiguration()
                    .withPrivateIpAddress(new PrivateIpAddress().withIpAddress("10.0.1.112").withSubnetResourceId(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/alternate"))
                    .withSqlVirtualMachineInstance(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm1")))
            .withPort(1433).create();
    }
}
