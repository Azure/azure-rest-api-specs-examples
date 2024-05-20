
import com.azure.resourcemanager.nginx.models.NginxConfigurationFile;
import com.azure.resourcemanager.nginx.models.NginxConfigurationPackage;
import com.azure.resourcemanager.nginx.models.NginxConfigurationProperties;
import java.util.Arrays;

/**
 * Samples for Configurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/
     * Configurations_CreateOrUpdate.json
     */
    /**
     * Sample code: Configurations_CreateOrUpdate.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void configurationsCreateOrUpdate(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.configurations().define("default").withExistingNginxDeployment("myResourceGroup", "myDeployment")
            .withProperties(new NginxConfigurationProperties()
                .withFiles(Arrays.asList(
                    new NginxConfigurationFile().withContent("ABCDEF==").withVirtualPath("/etc/nginx/nginx.conf")))
                .withPackageProperty(new NginxConfigurationPackage()).withRootFile("/etc/nginx/nginx.conf"))
            .create();
    }
}
