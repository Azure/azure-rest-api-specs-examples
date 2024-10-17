
import com.azure.resourcemanager.appcontainers.models.JavaComponent;
import com.azure.resourcemanager.appcontainers.models.JavaComponentConfigurationProperty;
import com.azure.resourcemanager.appcontainers.models.JavaComponentPropertiesScale;
import com.azure.resourcemanager.appcontainers.models.JavaComponentServiceBind;
import com.azure.resourcemanager.appcontainers.models.SpringBootAdminComponent;
import java.util.Arrays;

/**
 * Samples for JavaComponents Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/
     * JavaComponents_Patch_ServiceBind.json
     */
    /**
     * Sample code: Patch Java Component with ServiceBinds.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        patchJavaComponentWithServiceBinds(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        JavaComponent resource = manager.javaComponents()
            .getWithResponse("examplerg", "myenvironment", "myjavacomponent", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new SpringBootAdminComponent()
            .withConfigurations(Arrays.asList(
                new JavaComponentConfigurationProperty().withPropertyName("spring.boot.admin.ui.enable-toasts")
                    .withValue("true"),
                new JavaComponentConfigurationProperty().withPropertyName("spring.boot.admin.monitor.status-interval")
                    .withValue("10000ms")))
            .withScale(new JavaComponentPropertiesScale().withMinReplicas(1).withMaxReplicas(1))
            .withServiceBinds(Arrays.asList(new JavaComponentServiceBind().withName("yellowcat").withServiceId(
                "/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/javaComponents/yellowcat"))))
            .apply();
    }
}
