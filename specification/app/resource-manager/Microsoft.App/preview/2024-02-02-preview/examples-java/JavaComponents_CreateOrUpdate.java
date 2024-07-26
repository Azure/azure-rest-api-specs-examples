
import com.azure.resourcemanager.appcontainers.models.JavaComponentConfigurationProperty;
import com.azure.resourcemanager.appcontainers.models.SpringBootAdminComponent;
import java.util.Arrays;

/**
 * Samples for JavaComponents CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * JavaComponents_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update Java Component.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateJavaComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.javaComponents().define("myjavacomponent").withExistingManagedEnvironment("examplerg", "myenvironment")
            .withProperties(new SpringBootAdminComponent().withConfigurations(Arrays.asList(
                new JavaComponentConfigurationProperty().withPropertyName("spring.boot.admin.ui.enable-toasts")
                    .withValue("true"),
                new JavaComponentConfigurationProperty().withPropertyName("spring.boot.admin.monitor.status-interval")
                    .withValue("10000ms"))))
            .create();
    }
}
