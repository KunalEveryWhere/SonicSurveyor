#! /bin/bash

# Run the get started turorial
# https://noisemodelling.readthedocs.io/en/latest/Get_Started_Tutorial.html

# Step 4: Upload files to database
# create (or load existing) database and load a shape file into the database
# ./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/noisemodelling/wps/Import_and_Export/Import_File.groovy -pathFile NoiseModelling_without_gui/resources/org/noise_planet/noisemodelling/wps/ground_type.shp
# ./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/noisemodelling/wps/Import_and_Export/Import_File.groovy -pathFile NoiseModelling_without_gui/resources/org/noise_planet/noisemodelling/wps/buildings.shp
# ./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/oisemodelling/wps/Import_and_Export/Import_File.groovy -pathFile NoiseModelling_without_gui/resources/org/noise_planet/noisemodelling/wps/receivers.shp
# ./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/noisemodelling/wps/Import_and_Export/Import_File.groovy -pathFile NoiseModelling_without_gui/resources/org/noise_planet/noisemodelling/wps/ROADS2.shp
# ./NoiseModelling_without_gui/bin/wps_scripts -w ./ -s NoiseModelling_without_gui/noisemodelling/wps/Import_and_Export/Import_File.groovy -pathFile NoiseModelling_without_gui/resources/org/noise_planet/noisemodelling/wps/dem.geojson

# Step 1: Clean the DB
command_output=$(./NM_4.0.0_WO_GUI/bin/wps_scripts -w ./ -s  NoiseModelling_without_gui/noisemodelling/wps/Database_Manager/Clean_Database.groovy -areYouSure true)
#command_output=$(ls)
echo "$command_output"
# echo "Hello My Man"

sudo chmod 755 ./NM_4.0.0_WO_GUI/bin/wps_scripts
sudo chmod +x test.sh
./test.sh