
import com.azure.resourcemanager.appplatform.fluent.models.DeploymentResourceInner;
import com.azure.resourcemanager.appplatform.models.CustomContainer;
import com.azure.resourcemanager.appplatform.models.CustomContainerUserSourceInfo;
import com.azure.resourcemanager.appplatform.models.DeploymentResourceProperties;
import com.azure.resourcemanager.appplatform.models.ImageRegistryCredential;
import java.util.Arrays;

/**
 * Samples for Deployments Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * Deployments_Update_CustomContainer.json
     */
    /**
     * Sample code: Deployments_Update_CustomContainer.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsUpdateCustomContainer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getDeployments()
            .update("myResourceGroup", "myservice", "myapp", "mydeployment",
                new DeploymentResourceInner()
                    .withProperties(
                        new DeploymentResourceProperties()
                            .withSource(new CustomContainerUserSourceInfo()
                                .withCustomContainer(new CustomContainer().withServer("mynewacr.azurecr.io")
                                    .withContainerImage("myNewContainerImage:v1").withCommand(Arrays.asList("/bin/sh"))
                                    .withArgs(Arrays.asList("-c", "while true; do echo hello; sleep 10;done"))
                                    .withImageRegistryCredential(new ImageRegistryCredential()
                                        .withUsername("myNewUsername").withPassword("fakeTokenPlaceholder"))))),
                com.azure.core.util.Context.NONE);
    }
}
