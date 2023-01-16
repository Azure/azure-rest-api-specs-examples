import com.azure.resourcemanager.delegatednetwork.models.ControllerDetails;
import com.azure.resourcemanager.delegatednetwork.models.SubnetDetails;

/** Samples for DelegatedSubnetService PutDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/putDelegatedSubnet.json
     */
    /**
     * Sample code: put delegated subnet.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void putDelegatedSubnet(com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager
            .delegatedSubnetServices()
            .define("delegated1")
            .withRegion("West US")
            .withExistingResourceGroup("TestRG")
            .withSubnetDetails(
                new SubnetDetails()
                    .withId(
                        "/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet"))
            .withControllerDetails(
                new ControllerDetails()
                    .withId(
                        "/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/controller/dnctestcontroller"))
            .create();
    }
}
