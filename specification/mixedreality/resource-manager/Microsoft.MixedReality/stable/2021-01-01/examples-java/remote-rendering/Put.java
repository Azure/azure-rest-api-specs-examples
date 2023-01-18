import com.azure.resourcemanager.mixedreality.models.Identity;
import com.azure.resourcemanager.mixedreality.models.ResourceIdentityType;

/** Samples for RemoteRenderingAccounts Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/remote-rendering/Put.json
     */
    /**
     * Sample code: Create remote rendering account.
     *
     * @param manager Entry point to MixedRealityManager.
     */
    public static void createRemoteRenderingAccount(
        com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager
            .remoteRenderingAccounts()
            .define("MyAccount")
            .withRegion("eastus2euap")
            .withExistingResourceGroup("MyResourceGroup")
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .create();
    }
}
