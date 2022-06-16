import com.azure.core.util.Context;
import com.azure.resourcemanager.azurestackhci.models.Cluster;
import com.azure.resourcemanager.azurestackhci.models.ClusterDesiredProperties;
import com.azure.resourcemanager.azurestackhci.models.DiagnosticLevel;
import com.azure.resourcemanager.azurestackhci.models.WindowsServerSubscription;
import java.util.HashMap;
import java.util.Map;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/UpdateCluster.json
     */
    /**
     * Sample code: Update cluster.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void updateCluster(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        Cluster resource =
            manager.clusters().getByResourceGroupWithResponse("test-rg", "myCluster", Context.NONE).getValue();
        resource
            .update()
            .withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withCloudManagementEndpoint("https://98294836-31be-4668-aeae-698667faf99b.waconazure.com")
            .withDesiredProperties(
                new ClusterDesiredProperties()
                    .withWindowsServerSubscription(WindowsServerSubscription.ENABLED)
                    .withDiagnosticLevel(DiagnosticLevel.BASIC))
            .apply();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
