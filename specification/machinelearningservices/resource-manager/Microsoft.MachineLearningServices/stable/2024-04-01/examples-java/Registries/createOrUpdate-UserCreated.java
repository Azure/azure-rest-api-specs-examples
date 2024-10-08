
import com.azure.resourcemanager.machinelearning.models.AcrDetails;
import com.azure.resourcemanager.machinelearning.models.ArmResourceId;
import com.azure.resourcemanager.machinelearning.models.EndpointServiceConnectionStatus;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentity;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.machinelearning.models.PrivateEndpointResource;
import com.azure.resourcemanager.machinelearning.models.RegistryPrivateEndpointConnection;
import com.azure.resourcemanager.machinelearning.models.RegistryPrivateLinkServiceConnectionState;
import com.azure.resourcemanager.machinelearning.models.RegistryRegionArmDetails;
import com.azure.resourcemanager.machinelearning.models.Sku;
import com.azure.resourcemanager.machinelearning.models.SkuTier;
import com.azure.resourcemanager.machinelearning.models.StorageAccountDetails;
import com.azure.resourcemanager.machinelearning.models.UserAssignedIdentity;
import com.azure.resourcemanager.machinelearning.models.UserCreatedAcrAccount;
import com.azure.resourcemanager.machinelearning.models.UserCreatedStorageAccount;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Registries CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registries/createOrUpdate-UserCreated.json
     */
    /**
     * Sample code: CreateOrUpdate Registry with user created accounts.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateRegistryWithUserCreatedAccounts(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registries().define("string").withRegion("string").withExistingResourceGroup("test-rg")
            .withTags(mapOf())
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.NONE)
                .withUserAssignedIdentities(mapOf("string", new UserAssignedIdentity())))
            .withKind("string")
            .withSku(new Sku().withName("string").withTier(SkuTier.FREE).withSize("string").withFamily("string")
                .withCapacity(1))
            .withDiscoveryUrl("string").withIntellectualPropertyPublisher("string")
            .withManagedResourceGroup(new ArmResourceId().withResourceId("string")).withMlFlowRegistryUri("string")
            .withRegistryPrivateEndpointConnections(Arrays.asList(new RegistryPrivateEndpointConnection()
                .withId("string").withLocation("string").withGroupIds(Arrays.asList("string"))
                .withPrivateEndpoint(new PrivateEndpointResource().withSubnetArmId("string"))
                .withRegistryPrivateLinkServiceConnectionState(
                    new RegistryPrivateLinkServiceConnectionState().withActionsRequired("string")
                        .withDescription("string").withStatus(EndpointServiceConnectionStatus.APPROVED))
                .withProvisioningState("string")))
            .withPublicNetworkAccess("string")
            .withRegionDetails(Arrays.asList(new RegistryRegionArmDetails()
                .withAcrDetails(Arrays.asList(new AcrDetails().withUserCreatedAcrAccount(
                    new UserCreatedAcrAccount().withArmResourceId(new ArmResourceId().withResourceId("string")))))
                .withLocation("string")
                .withStorageAccountDetails(Arrays.asList(new StorageAccountDetails().withUserCreatedStorageAccount(
                    new UserCreatedStorageAccount().withArmResourceId(new ArmResourceId().withResourceId("string")))))))
            .create();
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
