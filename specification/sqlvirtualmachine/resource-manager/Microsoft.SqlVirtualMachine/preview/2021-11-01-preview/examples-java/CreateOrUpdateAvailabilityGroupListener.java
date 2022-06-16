import com.azure.resourcemanager.sqlvirtualmachine.models.LoadBalancerConfiguration;
import com.azure.resourcemanager.sqlvirtualmachine.models.PrivateIpAddress;
import java.util.Arrays;

/** Samples for AvailabilityGroupListeners CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/CreateOrUpdateAvailabilityGroupListener.json
     */
    /**
     * Sample code: Creates or updates an availability group listener.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void createsOrUpdatesAnAvailabilityGroupListener(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager
            .availabilityGroupListeners()
            .define("agl-test")
            .withExistingSqlVirtualMachineGroup("testrg", "testvmgroup")
            .withAvailabilityGroupName("ag-test")
            .withLoadBalancerConfigurations(
                Arrays
                    .asList(
                        new LoadBalancerConfiguration()
                            .withPrivateIpAddress(
                                new PrivateIpAddress()
                                    .withIpAddress("10.1.0.112")
                                    .withSubnetResourceId(
                                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"))
                            .withLoadBalancerResourceId(
                                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb-test")
                            .withProbePort(59983)
                            .withSqlVirtualMachineInstances(
                                Arrays
                                    .asList(
                                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2",
                                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm3"))))
            .withPort(1433)
            .create();
    }
}
