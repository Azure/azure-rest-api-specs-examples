
import com.azure.resourcemanager.nginx.models.AnalysisCreate;
import com.azure.resourcemanager.nginx.models.AnalysisCreateConfig;
import com.azure.resourcemanager.nginx.models.NginxConfigurationFile;
import com.azure.resourcemanager.nginx.models.NginxConfigurationPackage;
import java.util.Arrays;

/**
 * Samples for Configurations Analysis.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/Configurations_Analysis.
     * json
     */
    /**
     * Sample code: Configurations_Analysis.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void configurationsAnalysis(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.configurations().analysisWithResponse("myResourceGroup", "myDeployment", "default",
            new AnalysisCreate().withConfig(new AnalysisCreateConfig().withRootFile("/etc/nginx/nginx.conf")
                .withFiles(Arrays.asList(
                    new NginxConfigurationFile().withContent("ABCDEF==").withVirtualPath("/etc/nginx/nginx.conf")))
                .withPackageProperty(new NginxConfigurationPackage())),
            com.azure.core.util.Context.NONE);
    }
}
