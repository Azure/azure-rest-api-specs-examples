
import com.azure.resourcemanager.appplatform.fluent.models.ConfigServerResourceInner;
import com.azure.resourcemanager.appplatform.models.ConfigServerGitProperty;
import com.azure.resourcemanager.appplatform.models.ConfigServerProperties;
import com.azure.resourcemanager.appplatform.models.ConfigServerSettings;
import java.util.Arrays;

/**
 * Samples for ConfigServers UpdatePatch.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ConfigServers_UpdatePatch.json
     */
    /**
     * Sample code: ConfigServers_UpdatePatch.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configServersUpdatePatch(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getConfigServers().updatePatch("myResourceGroup", "myservice",
            new ConfigServerResourceInner().withProperties(
                new ConfigServerProperties().withConfigServer(new ConfigServerSettings().withGitProperty(
                    new ConfigServerGitProperty().withUri("https://github.com/fake-user/fake-repository.git")
                        .withLabel("master").withSearchPaths(Arrays.asList("/"))))),
            com.azure.core.util.Context.NONE);
    }
}
