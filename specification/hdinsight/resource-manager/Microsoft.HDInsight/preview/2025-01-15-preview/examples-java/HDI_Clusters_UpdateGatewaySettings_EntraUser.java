
import com.azure.resourcemanager.hdinsight.models.EntraUserInfo;
import com.azure.resourcemanager.hdinsight.models.UpdateGatewaySettingsParameters;
import java.util.Arrays;

/**
 * Samples for Clusters UpdateGatewaySettings.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2025-01-15-preview/examples/
     * HDI_Clusters_UpdateGatewaySettings_EntraUser.json
     */
    /**
     * Sample code: Update Entra User In HDInsight.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void updateEntraUserInHDInsight(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.clusters().updateGatewaySettings("rg1", "cluster1",
            new UpdateGatewaySettingsParameters().withRestAuthEntraUsers(
                Arrays.asList(new EntraUserInfo().withObjectId("00000000-0000-0000-0000-000000000000")
                    .withDisplayName("displayName").withUpn("user@microsoft.com"))),
            com.azure.core.util.Context.NONE);
    }
}
