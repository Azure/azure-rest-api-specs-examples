
import com.azure.resourcemanager.confluent.models.KeyValuePair;
import com.azure.resourcemanager.confluent.models.NetworkGatewayProperties;
import java.util.Arrays;

/**
 * Samples for NetworkGatewayResources CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/NetworkGatewayResources_CreateOrReplace_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkGatewayResources_CreateOrReplace_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        networkGatewayResourcesCreateOrReplaceMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.networkGatewayResources().define("gw-def456")
            .withExistingEnvironment("rgconfluent", "myOrganization", "env-abc123")
            .withProperties(
                new NetworkGatewayProperties().withNetworkGatewayName("my-network-gateway").withRegion("eastus")
                    .withDictionary(Arrays.asList(
                        new KeyValuePair().withKey("fakeTokenPlaceholder").withValue("PrivateNetworkInterface"))))
            .create();
    }
}
