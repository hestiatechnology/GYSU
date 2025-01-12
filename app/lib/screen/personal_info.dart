import 'package:flutter/material.dart';

class PersonalInfo extends StatefulWidget {
  const PersonalInfo({super.key});

  @override
  PersonalInfoState createState() => PersonalInfoState();
}

class PersonalInfoState extends State<PersonalInfo> {
  final _formKey = GlobalKey<FormState>();
  String _name = '';
  String _description = '';
  String _university = 'Politécnico do Cávado e do Ave'; // Default value
  String _course = '';
  final List<String> _selectedInterests = []; // To store selected interests

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Registo - Dados Pessoais'),
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Form(
            key: _formKey,
            child: Column(
              children: <Widget>[
                Center(
                  child: CircleAvatar(
                    radius: 50,
                    child: Icon(Icons.person, size: 50),
                  ),
                ),
                SizedBox(height: 20),
                TextFormField(
                  decoration: InputDecoration(
                    labelText: 'Nome',
                    hintText: 'Digite seu nome',
                    border: OutlineInputBorder(),
                  ),
                  validator: (value) {
                    if (value == null || value.isEmpty) {
                      return 'Por favor, insira seu nome';
                    }
                    return null;
                  },
                  onSaved: (value) {
                    _name = value!;
                  },
                ),
                SizedBox(height: 20),
                TextFormField(
                  maxLines: 4,
                  decoration: InputDecoration(
                    labelText: 'Descrição',
                    hintText: 'Fale um pouco sobre você',
                    border: OutlineInputBorder(),
                  ),
                  validator: (value) {
                    if (value == null || value.isEmpty) {
                      return 'Por favor, insira uma descrição';
                    }
                    return null;
                  },
                  onSaved: (value) {
                    _description = value!;
                  },
                ),
                SizedBox(height: 20),
                Row(
                  children: [
                    Text('Sou estudante'),
                    Spacer(),
                    Switch(
                      value: true, // Initially checked
                      onChanged: (value) {
                        // Handle switch change if needed
                      },
                    ),
                  ],
                ),
                SizedBox(height: 20),
                TextFormField(
                  decoration: InputDecoration(
                    labelText: 'Universidade',
                    hintText: 'Digite sua universidade',
                    border: OutlineInputBorder(),
                  ),
                  //validator: (value) {
                  //  if (value == null || value.isEmpty) {
                  //    return 'Por favor, insira sua universidade';
                  //  }
                  //  return null;
                  //},
                  onSaved: (value) {
                    _university = value!;
                  },
                ),
                SizedBox(height: 20),
                Row(
                  children: [
                    Expanded(
                      child: TextFormField(
                        decoration: InputDecoration(
                          labelText: 'Curso',
                          hintText: 'Digite seu curso',
                          border: OutlineInputBorder(),
                        ),
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'Por favor, insira seu curso';
                          }
                          return null;
                        },
                        onSaved: (value) {
                          _course = value!;
                        },
                      ),
                    ),
                    SizedBox(width: 20),
                    SizedBox(
                      width: 150, // Set the desired width
                      child: DropdownButtonFormField<int>(
                        decoration: InputDecoration(
                          labelText: 'Ano',
                          border: OutlineInputBorder(),
                        ),
                        value: 1, // Default value
                        items: List.generate(5, (index) {
                          return DropdownMenuItem<int>(
                            value: index + 1,
                            child: Text('${index + 1}º Ano'),
                          );
                        }),
                        onChanged: (value) {
                          // Handle dropdown change if needed
                        },
                      ),
                    ),
                  ],
                ),
                SizedBox(height: 20),
                Text('Interesses:'),
                Wrap(
                  spacing: 8.0,
                  runSpacing: 4.0,
                  children: [
                    ChoiceChip(
                      label: Text('Informática'),
                      selected: _selectedInterests.contains('Informática'),
                      onSelected: (bool selected) {
                        setState(() {
                          if (selected) {
                            _selectedInterests.add('Informática');
                          } else {
                            _selectedInterests.remove('Informática');
                          }
                        });
                      },
                    ),
                    ChoiceChip(
                      label: Text('Marketing'),
                      selected: _selectedInterests.contains('Marketing'),
                      onSelected: (bool selected) {
                        setState(() {
                          if (selected) {
                            _selectedInterests.add('Marketing');
                          } else {
                            _selectedInterests.remove('Marketing');
                          }
                        });
                      },
                    ),
                    ChoiceChip(
                      label: Text('Gestão'),
                      selected: _selectedInterests.contains('Gestão'),
                      onSelected: (bool selected) {
                        setState(() {
                          if (selected) {
                            _selectedInterests.add('Gestão');
                          } else {
                            _selectedInterests.remove('Gestão');
                          }
                        });
                      },
                    ),
                    ChoiceChip(
                      label: Text('Recursos Humanos'),
                      selected: _selectedInterests.contains('Recursos Humanos'),
                      onSelected: (bool selected) {
                        setState(() {
                          if (selected) {
                            _selectedInterests.add('Recursos Humanos');
                          } else {
                            _selectedInterests.remove('Recursos Humanos');
                          }
                        });
                      },
                    ),
                  ],
                ),
                SizedBox(height: 30),
                FilledButton(
                  onPressed: () {
                    if (_formKey.currentState!.validate()) {
                      _formKey.currentState!.save();
                      // Handle registration logic here
                      // e.g., send data to server, navigate to next screen
                    }
                  },
                  child: Text('Guardar'),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
