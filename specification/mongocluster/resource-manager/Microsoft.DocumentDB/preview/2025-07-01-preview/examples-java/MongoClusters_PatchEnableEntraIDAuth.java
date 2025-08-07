
import com.azure.resourcemanager.mongocluster.models.AuthConfigProperties;
import com.azure.resourcemanager.mongocluster.models.AuthenticationMode;
import com.azure.resourcemanager.mongocluster.models.MongoCluster;
import com.azure.resourcemanager.mongocluster.models.MongoClusterUpdateProperties;
import java.util.Arrays;

/**
 * Samples for MongoClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/MongoClusters_PatchEnableEntraIDAuth.json
     */
    /**
     * Sample code: Updates the allowed authentication modes to include Microsoft Entra ID authentication.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void updatesTheAllowedAuthenticationModesToIncludeMicrosoftEntraIDAuthentication(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        MongoCluster resource = manager.mongoClusters()
            .getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new MongoClusterUpdateProperties().withAuthConfig(new AuthConfigProperties()
            .withAllowedModes(Arrays.asList(AuthenticationMode.NATIVE_AUTH, AuthenticationMode.MICROSOFT_ENTRA_ID))))
            .apply();
    }
}
