
import com.azure.resourcemanager.network.fluent.models.VerifierWorkspaceInner;
import com.azure.resourcemanager.network.models.VerifierWorkspaceProperties;

/**
 * Samples for VerifierWorkspaces Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VerifierWorkspacePut.json
     */
    /**
     * Sample code: VerifierWorkspaceCreate.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void verifierWorkspaceCreate(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVerifierWorkspaces().createWithResponse("rg1", "testNetworkManager", "testWorkspace",
            new VerifierWorkspaceInner().withLocation("eastus")
                .withProperties(new VerifierWorkspaceProperties().withDescription("A sample workspace")),
            null, com.azure.core.util.Context.NONE);
    }
}
