import com.azure.core.util.Context;
import com.azure.resourcemanager.hdinsight.models.UpdateClusterIdentityCertificateParameters;

/** Samples for Clusters UpdateIdentityCertificate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Clusters_UpdateClusterIdentityCertificate.json
     */
    /**
     * Sample code: Update cluster identity certificate.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void updateClusterIdentityCertificate(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .clusters()
            .updateIdentityCertificate(
                "rg1",
                "cluster1",
                new UpdateClusterIdentityCertificateParameters()
                    .withApplicationId("applicationId")
                    .withCertificate("base64encodedcertificate")
                    .withCertificatePassword("**********"),
                Context.NONE);
    }
}
