import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.SshPublicKeyResourceInner;

/** Samples for SshPublicKeys Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateAnSshPublicKey.json
     */
    /**
     * Sample code: Create a new SSH public key resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createANewSSHPublicKeyResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSshPublicKeys()
            .createWithResponse(
                "myResourceGroup",
                "mySshPublicKeyName",
                new SshPublicKeyResourceInner().withLocation("westus").withPublicKey("{ssh-rsa public key}"),
                Context.NONE);
    }
}
